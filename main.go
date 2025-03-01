package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		if len(args) == 1 {
			fmt.Println("no website provided")
			os.Exit(1)
		} else {
			fmt.Println("too many arguments provided")
			os.Exit(1)
		}
	}

	rawBaseURL := args[1]

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	var pages = make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)
	fmt.Printf("Total crawled pages: %d\n", len(pages))
}