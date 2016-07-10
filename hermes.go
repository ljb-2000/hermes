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
	supervisor.Register("XKCD update checker", &xkcd)
	go supervisor.Run()

	RunWebServer(&supervisor)
}
