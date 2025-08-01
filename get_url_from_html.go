package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTLM(htmlBody, rawBasedURL string) ([]string, error) {
	url, err := normalizeURL(rawBasedURL)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", url)
	parsedHTML, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", parsedHTML)
	fmt.Println(parsedHTML.Type)
	fmt.Printf("%+v\n", parsedHTML.ChildNodes())
	return nil, nil
}
