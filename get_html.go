package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("netowrk error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("http error: %s", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-html response: %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(body), nil
}