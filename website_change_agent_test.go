package main

import (
	"io/ioutil"
	"testing"
)

func TestFetcherReaderIsClosed(t *testing.T) {
	reader := CloseChecker{}
	fetcher := StubFetcher{&reader}

	FetchWebsiteChecksum(fetcher, "some url")

	if !reader.IsClosed {
		t.Error("the reader was not closed")
	}
}

func TestChecksumsAreIdenticalWithSameContent(t *testing.T) {
	content := "some content"

	fetchOne := NewStubFetcher(content)
	fetchTwo := NewStubFetcher(content)

	checkOne, _ := FetchWebsiteChecksum(fetchOne, "some url")
	checkTwo, _ := FetchWebsiteChecksum(fetchTwo, "some url")

	if checkOne != checkTwo {
		t.Errorf("%v != %v\n", checkOne, checkTwo)
	}
}

func TestChecksumIsZeroWhenFetchErrorOccurs(t *testing.T) {
	fetcher := BrokenFetcher{}

	size, _ := FetchWebsiteChecksum(fetcher, "some url")

	if size != 0 {
		t.Errorf("expected checksum to be 0, got %v\n", size)
	}
}

func TestChecksumIsZeroWhenReadErrorOccurs(t *testing.T) {
	reader := ioutil.NopCloser(BrokenReader{})
	fetcher := StubFetcher{reader}

	size, _ := FetchWebsiteChecksum(fetcher, "some url")

	if size != 0 {
		t.Errorf("expected checksum to be 0, got %v\n", size)
	}
}

func TestAgentMakesARequestWhenStarted() {
	fetcher := NewStubFetcher("hello there")
	agent := StartWebsiteChangeAgent(fetcher)

	select {
	case <-agent.Events:

	}
}
