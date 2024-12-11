package main

import (
	"errors"
	"net/url"
	"strings"
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