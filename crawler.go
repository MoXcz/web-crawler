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

	normBaseURL, err := normalizeURL(rawBaseURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg.addPageVisit(normBaseURL)

	for _, link := range links {
		normLink, err := normalizeURL(link)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if !cfg.addPageVisit(normLink) {
			continue
		}

		cfg.wg.Add(1)
		go func(link string) {
			cfg.concurrencyControl <- struct{}{}
			defer cfg.wg.Done()
			defer func() { <-cfg.concurrencyControl }()
			cfg.crawlPage(link)
		}(link)
	}
}

func compareHostURLs(rawBaseURL, rawCurrentURL string) (bool, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return false, fmt.Errorf("could not parse URL string")
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false, fmt.Errorf("could not parse URL string")
	}

	if baseURL.Hostname() != currentURL.Hostname() {
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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}
