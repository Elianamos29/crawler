package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
	REPORT for %s
=============================`, baseURL)

	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", pages[page], page)
	}
}

func sortPages(pages map[string]int) []string {
	sortedPages := make([]string, 0, len(pages))
	for normalizedURL := range pages {
		sortedPages = append(sortedPages, normalizedURL)
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		if pages[sortedPages[i]] == pages[sortedPages[j]] {
			return sortedPages[i] < sortedPages[j]
		}
		return pages[sortedPages[i]] > pages[sortedPages[j]]
	})
	return sortedPages
}
