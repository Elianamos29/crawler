# Web Crawler

Boot.dev challenge: building a web crawler using Go.

## Description

The web crawler traverses websites, collecting and processing data from web pages. It demonstrates the use of Go's concurrency features and HTTP handling capabilities.

## Features

- Crawl web pages starting from a given URL.
- Extract and display links from the crawled pages.
- Handle concurrency for efficient crawling.

## Requirements

- Go 1.20 or later

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Elianamos29/crawler.git
    cd crawler
    ```

2. Build the project:
    ```bash
    go build
    ```

## Usage

Run the crawler with a starting URL:
```bash
./crawler https://example.com
```