package main

import (
	"fmt"
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
	pages := map[string]int{}
	fmt.Printf("starting crawl of: %s\n", argPage)
	crawlPage(argPage, argPage, pages)

	for page, count := range pages {
		fmt.Printf("%d - %s\n", count, page)
	}
}
