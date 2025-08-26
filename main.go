package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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
	ch := make(chan struct{}, 5)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	cfg := config{
		pages:              pages,
		baseURL:            parsedURL,
		concurrencyControl: ch,
		wg:                 wg,
		mu:                 mu,
	}

	fmt.Printf("starting crawl of: %s\n", argPage)

	wg.Add(1)
	go cfg.crawlPage(argPage)

	wg.Wait()
	defer close(ch)

	for page, count := range pages {
		fmt.Printf("%d - %s\n", count, page)
	}
}
