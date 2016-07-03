package main

import "time"

func main() {
	fetcher := HTTPFetcher{}

	supervisor := NewAgentSupervisor()
	xkcd := NewWebsiteChangeAgent(WebsiteChangeAgentConfig{
		Fetcher:  fetcher,
		Name:     "http://www.xkcd.com",
		Interval: time.Minute,
	})
	supervisor.Register(&xkcd)
	supervisor.Run()

	RunWebServer(&supervisor)
}
