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

	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}