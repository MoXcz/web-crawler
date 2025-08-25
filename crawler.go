package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	rawBaseURL := cfg.baseURL.String()
	eq, err := compareHostURLs(rawBaseURL, rawCurrentURL)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if !eq {
		log.Printf("%s has not the same domain as %s\n", rawBaseURL, rawCurrentURL)
		return
	}

	normCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, ok := cfg.pages[normCurrentURL]; ok {
		cfg.pages[normCurrentURL]++
		return
	}
	cfg.pages[normCurrentURL] = 1
	fmt.Printf("Crawling page: %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not get HTML for %s: %v\n", rawCurrentURL, err)
		return
	}

	links, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("could not get links for %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, link := range links {
		cfg.crawlPage(link)
	}
}

func compareHostURLs(rawBaseURL, rawCurrentURL string) (bool, error) {
	basedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return false, fmt.Errorf("could not parse URL string")
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false, fmt.Errorf("could not parse URL string")
	}

	if basedURL.Hostname() != currentURL.Hostname() {
		return false, nil
	}
	return true, nil
}

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error response: %s", resp.Body)
	}
	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	if contentType != "text/html" {
		return "", fmt.Errorf("not valid Content-Type: %s", resp.Header.Get("Content-Type"))
	}

	htmlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(htmlContent), nil
}
