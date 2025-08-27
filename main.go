package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 4 {
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "not enough arguments\n\n")
		}
		if len(os.Args) > 4 {
			fmt.Fprintf(os.Stderr, "too many arguments provided\n\n")
		}
		usage()

		os.Exit(1)
	}

	argPage := os.Args[1]
	argMaxConcurrency := os.Args[2]
	argMaxPages := os.Args[3]

	maxConcurrency, err := strconv.Atoi(argMaxConcurrency)
	if err != nil {
		fmt.Println("Invalid value: Concurrency Control argument should be a number")
		usage()
		return
	}
	maxPages, err := strconv.Atoi(argMaxPages)
	if err != nil {
		fmt.Println("Invalid value: Max Pages argument should be a number")
		usage()
		return
	}
	fmt.Printf("maxPages %d, maxConcurrency: %d", maxPages, maxConcurrency)

	parsedURL, err := url.Parse(argPage)
	if err != nil {
		return
	}
	pages := map[string]int{}
	ch := make(chan struct{}, maxConcurrency)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	cfg := config{
		pages:              pages,
		baseURL:            parsedURL,
		concurrencyControl: ch,
		wg:                 wg,
		mu:                 mu,
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", argPage)

	cfg.wg.Add(1)
	go cfg.crawlPage(argPage)

	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseURL.String())
}

func usage() {
	fmt.Fprintf(os.Stdin, "Usage: crawler [URL] [Concurrency Control] [Max Pages]\n")
	fmt.Fprintf(os.Stdin, "\tURL: Base URL to start crawling\n")
	fmt.Fprintf(os.Stdin, "\tConcurrency Control: How many goroutines to run simultaneously\n")
	fmt.Fprintf(os.Stdin, "\tMax Pages: Max number of pages to crawl through\n")
}
