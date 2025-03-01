package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	ok, err := sameDomain(rawBaseURL, rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !ok {
		return
	}

	normalizedCurrentUrL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normaizing url: %v", err)
		return
	}

	if _, exists := pages[normalizedCurrentUrL]; exists {
		pages[normalizedCurrentUrL]++
		return
	}

	pages[normalizedCurrentUrL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v", err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}

func sameDomain(baseURL, refURL string) (bool, error) {
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return false, fmt.Errorf("couldn't parse url %s: %v", baseURL, err)
	}

	parsedRef, err := url.Parse(refURL)
	if err != nil {
		return false, fmt.Errorf("couldn't parse url %s: %v", refURL, err)
	}

	baseHost := strings.TrimPrefix(parsedBase.Host, "www.")
	refHost := strings.TrimPrefix(parsedRef.Host, "www.")

	return baseHost == refHost, nil
}