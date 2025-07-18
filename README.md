# Go XPath HTML Scraper

A lightweight web scraping tool written in Go that retrieves HTML elements based on simplified XPath-like expressions. This project demonstrates how to parse and traverse HTML documents using Go's [`golang.org/x/net/html`](https://pkg.go.dev/golang.org/x/net/html) package, allowing you to extract specific elements from a web page.

---

## Features

- Parse and navigate HTML documents
- Simplified XPath support (basic node name + positional predicates like `/div[2]/ul/li[3]`)
- Extract elements and their text content
- Fetch and parse remote HTML pages via HTTP

---

## Example

The included example fetches the contents of [scrapeme.live/shop](https://scrapeme.live/shop) and retrieves the third product's title using a path like:

```
/html/body/div[1]/div[2]/div/div[2]/main/ul/li[3]/a[1]/h2
```

---

## Dependencies

- [`golang.org/x/net/html`](https://pkg.go.dev/golang.org/x/net/html)

---

## Usage

```bash
go run main.go
```
