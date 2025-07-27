package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(URLString string) (string, error) {
	unparsedURL, err := url.Parse(URLString)
	if err != nil {
		return "", fmt.Errorf("could not parse URL string")
	}
	// Remove trailing "/" and " " from lowercase host/path
	parsedURL := strings.ToLower(strings.TrimRight(unparsedURL.Host+unparsedURL.Path, "/ "))
	return parsedURL, nil
}

func getURLsFromHTLM(htmlBody, rawBasedURL string) ([]string, error) {
	return nil, nil
}
