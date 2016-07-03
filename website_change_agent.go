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
	fetcher Fetcher
	events  chan bool
	name    string
	lastSum int64
}

func (w WebsiteChangeAgent) Events() chan bool {
	return w.events
}

func (w *WebsiteChangeAgent) HasContentChanged() bool {
	sum, err := FetchWebsiteChecksum(w.fetcher, w.name)
	if err != nil {
		panic("uh oh")
	}
	if sum != w.lastSum {
		w.lastSum = sum
		return true
	}
	return false
}

func (w *WebsiteChangeAgent) run() {
	for {
		if w.HasContentChanged() {
			w.events <- true
		}
	}
}

func NewWebsiteChangeAgent(fetcher Fetcher, name string) WebsiteChangeAgent {
	events := make(chan bool)
	agent := WebsiteChangeAgent{fetcher, events, name}
	go agent.run()
	return agent
}
