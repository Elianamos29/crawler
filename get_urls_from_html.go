package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var urls []string
	uniqueURLs := make(map[string]bool)

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	extractURLs(doc, baseURL, uniqueURLs)

	for url := range uniqueURLs {
		urls = append(urls, url)
	}

	return urls, nil
}

func extractURLs(n *html.Node, baseURL *url.URL, urls map[string]bool) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				resolvedURL := resolveURL(attr.Val, baseURL)
				if resolvedURL != "" {
					urls[resolvedURL] = true
				}
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractURLs(c, baseURL, urls)
	}
}

func resolveURL(href string, baseURL *url.URL) string {
	parsedHref, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolvedURL := baseURL.ResolveReference(parsedHref)
	return resolvedURL.String()
}