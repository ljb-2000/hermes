package main

import (
	"io"
	"net/http"
)

// HTTPFetcher implements the Fetcher interface for URLs
type HTTPFetcher struct{}

// Fetch takes a URL, and returns a io.ReadCloser to the content returned
// by a GET request to that URL
func (h HTTPFetcher) Fetch(url string) (io.ReadCloser, error) {
	site, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return site.Body, nil
}
