package main

import (
	"hash/fnv"
	"io"
)

// FetchWebsiteChecksum downloads the resource at url,
// and returns an fnv hash of its content.
func FetchWebsiteChecksum(FetchWebsiteChecksum Fetcher, url string) (uint64, error) {
	website, err := FetchWebsiteChecksum.Fetch(url)
	if err != nil {
		return 0, err
	}
	defer website.Close()

	hash := fnv.New64()
	_, err = io.Copy(hash, website)
	if err != nil {
		return 0, err
	}

	return hash.Sum64(), nil
}

// WebsiteChangeAgent records when the content of a
// website has changed.
type WebsiteChangeAgent struct {
	lastChecksum uint64
}
