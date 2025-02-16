package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(inputURL string) (string, error) {
	if strings.TrimSpace(inputURL) == "" {
		return "", errors.New("input URL is empty")
	}

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Path == "" {
		parsedURL.Path = "/"
	}

	return parsedURL.Host + parsedURL.Path, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	uniqueURLs := make(map[string]bool)

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	extractURLs(doc, rawBaseURL, uniqueURLs)

	for url := range uniqueURLs {
		urls = append(urls, url)
	}

	return urls, nil
}

func extractURLs(n *html.Node, baseURL string, urls map[string]bool) {
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

func resolveURL(href, baseURL string) string {
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	parsedHref, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolvedURL := parsedBase.ResolveReference(parsedHref)
	return resolvedURL.String()
}