# web-crawler

> Made thanks to [Boot.dev](https://www.boot.dev)

This is an implementation of a Web Crawler using to Go programming language. It uses a customizable number of Goroutines to go through the crawled pages.

## Quick Start

```sh
go install github.com/MoXcz/web-crawler
web-crawler [URL] [Goroutines] [Pages]
```

- `URL` the base URL to crawl
- `Goroutines` the max number of Goroutines to use
- `Pages` the max number of pages to go through

## Installation

```sh
git clone https://github.com/MoXcz/web-crawler
cd web-crawler
go run .
```

Or as I prefer it:

```sh
git clone https://github.com/MoXcz/web-crawler
cd web-crawler
make run
make test
```

