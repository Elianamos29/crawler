package main

import (
	"fmt"
	"log"
)

func main() {
	htmlContent := `<html>
		<head><title>Test</title></head>
		<body>
			<a href="https://example.com">Example</a>
			<a href="/relative-link">Relative</a>
			<p>Some text <a href="https://another.com">Another</a></p>
		</body>
	</html>`
	urls, err := getURLsFromHTML(htmlContent, "")
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range urls {
		fmt.Println(url)
	}
}