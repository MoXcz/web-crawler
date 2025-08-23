package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBasedURL string) ([]string, error) {
	parsedHTML, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	links := cleanURLs(getURLs(parsedHTML), rawBasedURL)
	return links, nil
}

func getURLFromNode(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func getURLs(root *html.Node) []string {
	var links []string
	next := root
	for {
		if next != nil {
			if next.Data == "a" {
				links = append(links, getURLFromNode(next))
			}
		} else {
			break
		}

		if next.FirstChild != nil {
			links = append(links, getURLs(next.FirstChild)...)
		}
		next = next.NextSibling
	}

	return links
}

func cleanURLs(links []string, rawBasedURL string) []string {
	var cleanLinks []string
	for _, link := range links {
		if link[0] == '/' {
			link = rawBasedURL + link
		}
		cleanLinks = append(cleanLinks, strings.Trim(link, " "))
	}
	return cleanLinks
}
