package main

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type CloseChecker struct {
	IsClosed bool
}

func (c CloseChecker) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (c *CloseChecker) Close() error {
	c.IsClosed = true
	return nil
}

type StubFetcher struct {
	reader io.ReadCloser
}

func (s StubFetcher) Fetch(path string) (io.ReadCloser, error) {
	return s.reader, nil
}

type BrokenReader struct{}

func (b BrokenReader) Read([]byte) (int, error) {
	return 0, errors.New("read on broken reader")
}

type BrokenFetcher struct{}

func (b BrokenFetcher) Fetch(path string) (io.ReadCloser, error) {
	return nil, errors.New("fetch on broken fetcher")
}

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

	readerOne := ioutil.NopCloser(strings.NewReader(content))
	fetcherOne := StubFetcher{readerOne}

	readerTwo := ioutil.NopCloser(strings.NewReader(content))
	fetcherTwo := StubFetcher{readerTwo}

	checkOne, _ := FetchWebsiteChecksum(fetcherOne, "some url")
	checkTwo, _ := FetchWebsiteChecksum(fetcherTwo, "some url")

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
