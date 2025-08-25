package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) < 2 {
			fmt.Fprintf(os.Stderr, "no website provided\n\n")
		}
		if len(os.Args) > 2 {
			fmt.Fprintf(os.Stderr, "too many arguments provided\n\n")
		}
		fmt.Fprintf(os.Stderr, "Usage: crawler [URL]\n")
		os.Exit(1)
	}

	argPage := os.Args[1]
	parsedURL, err := url.Parse(argPage)
	if err != nil {
		return
	}
	pages := map[string]int{}
	cfg := config{
		pages:   pages,
		baseURL: parsedURL,
	}
	fmt.Printf("starting crawl of: %s\n", argPage)
	cfg.crawlPage(argPage)

	for page, count := range pages {
		fmt.Printf("%d - %s\n", count, page)
	}
}
