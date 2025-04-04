package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}
	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := args[1]
	maxConcurrency, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - invalid value for maxConcurrency: %v", err)
		return
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Error - invalid value for maxPages: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}
