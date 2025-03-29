package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedCurrentUrL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normaizing url: %v", err)
		return
	}

	if isFirst := cfg.addPageVisit(normalizedCurrentUrL); !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}