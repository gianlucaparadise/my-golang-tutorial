package main

import "testing"

func (safeFetcher *SafeFetcher) GetFetched() map[string]bool {
	return safeFetcher.fetched
}

func TestFetch(t *testing.T) {
	expectedUrls := [...]string{
		"https://golang.org/",
		"https://golang.org/pkg/",
		"https://golang.org/cmd/",
		"https://golang.org/pkg/fmt/",
		"https://golang.org/pkg/os/",
		"https://golang.org/pkg/pippo/",
	}

	c := SafeFetcher{fetched: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &c)

	actualUrls := c.GetFetched()

	for _, expected := range expectedUrls {
		_, matches := actualUrls[expected]
		if !matches {
			t.Fatalf(`Missing url in fetched: Expected = %v   Actual = %v`, expected, actualUrls)
		}
	}

	if len(actualUrls) != len(expectedUrls) {
		t.Fatalf(`Expected and Actual sizes differ: Expected = %v   Actual = %v`, expectedUrls, actualUrls)
	}
}
