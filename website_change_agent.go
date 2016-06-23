package main

import (
	"hash/fnv"
	"io"
)

// FetchWebsiteChecksum downloads the resource at url,
// and returns an fnv hash of its content.
func FetchWebsiteChecksum(fetcher Fetcher, url string) (uint64, error) {
	website, err := fetcher.Fetch(url)
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
	events chan bool
}

func (w WebsiteChangeAgent) Events() chan bool {
	return w.events
}

func (w *WebsiteChangeAgent) run() {
	w.events <- true
}

func NewWebsiteChangeAgent(fetcher Fetcher) WebsiteChangeAgent {
	events := make(chan bool)
	agent := WebsiteChangeAgent{events}
	go agent.run()
	return agent
}
