package main

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestFetcherReaderIsClosed(t *testing.T) {
	reader := CloseChecker{}
	fetcher := StubFetcher{&reader}

	FetchWebsiteChecksum(&fetcher, "some url")

	if !reader.IsClosed {
		t.Error("the reader was not closed")
	}
}

func TestChecksumsAreIdenticalWithSameContent(t *testing.T) {
	content := "some content"

	fetchOne := NewStringFetcher(content)
	fetchTwo := NewStringFetcher(content)

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

	size, _ := FetchWebsiteChecksum(&fetcher, "some url")

	if size != 0 {
		t.Errorf("expected checksum to be 0, got %v\n", size)
	}
}

func TestAgentMakesARequestWhenStarted(t *testing.T) {
	recorder := NewFetchRecorder(NewStringFetcher("some content"))
	agent := NewWebsiteChangeAgent(&recorder, "some name")

	select {
	case <-agent.Events():
	case <-time.After(1 * time.Second):
		t.Error("no event was sent at startup")
	}

	if len(recorder.Fetches) != 1 {
		t.Error("no fetch was made, but event was sent")
	}
}

func TestAgentSendsARequestWhenContentChanges(t *testing.T) {
	fetcher := NewStringFetcher("initial content")
	recorder := NewFetchRecorder(fetcher)
	agent := NewWebsiteChangeAgent(&recorder, "some name")

	select {
	case <-agent.Events():
	case <-time.After(1 * time.Second):
		t.Error("no initial event was sent")
	}

	fetcher.SetContent("new content")

	select {
	case <-agent.Events():
	case <-time.After(1 * time.Second):
		t.Error("no event was sent after change")
	}
}

func TestAgentSendsNothingWhenNothingChanges(t *testing.T) {
	select {}
}
