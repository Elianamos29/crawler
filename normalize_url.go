package main

import (
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host + parsedURL.Path, nil
}