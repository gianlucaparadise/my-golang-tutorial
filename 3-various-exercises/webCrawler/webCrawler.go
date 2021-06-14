package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeFetcher is safe to use concurrently.
type SafeFetcher struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func (safeFetcher *SafeFetcher) SaveUrlIfNeeded(url string) bool {
	// Checking the map
	safeFetcher.mu.Lock()
	defer safeFetcher.mu.Unlock()

	_, isFetched := safeFetcher.fetched[url]
	if isFetched {
		return false // This url has already been processed
	}

	safeFetcher.fetched[url] = true
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, safeFetcher *SafeFetcher) {
	if depth <= 0 {
		return
	}

	needsFetch := safeFetcher.SaveUrlIfNeeded(url)
	if !needsFetch {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, safeFetcher)
	}
	return
}

func main() {
	c := SafeFetcher{fetched: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &c)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
