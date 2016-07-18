package main

import "time"

func main() {
	fetcher := HTTPFetcher{}

	supervisor := NewAgentSupervisor()
	reddit := NewWebsiteChangeAgent(WebsiteChangeAgentConfig{
		Fetcher:  fetcher,
		Name:     "http://www.reddit.com",
		Interval: time.Minute,
	})
	xkcd := NewWebsiteChangeAgent(WebsiteChangeAgentConfig{
		Fetcher:  fetcher,
		Name:     "http://www.xkcd.com",
		Interval: time.Minute,
	})
	supervisor.Register("XKCD update checker", &xkcd)
	supervisor.Register("Reddit updates", &reddit)
	go supervisor.Run()

	RunWebServer(&supervisor)
}
