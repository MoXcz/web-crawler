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

	fmt.Printf("starting crawl of: %s\n", os.Args[1])
}
