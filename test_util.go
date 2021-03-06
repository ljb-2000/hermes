package main

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
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

type StringFetcher struct {
	content string
}

func (s *StringFetcher) SetContent(content string) {
	s.content = content
}

func (s StringFetcher) Fetch(path string) (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader(s.content)), nil
}

func NewStringFetcher(content string) StringFetcher {
	return StringFetcher{content}
}

type StubFetcher struct {
	Reader io.ReadCloser
}

func (s *StubFetcher) Fetch(name string) (io.ReadCloser, error) {
	return s.Reader, nil
}

type FetchRecorder struct {
	Fetcher
	Fetches []string
}

func (r *FetchRecorder) Fetch(path string) (io.ReadCloser, error) {
	r.Fetches = append(r.Fetches, path)
	return r.Fetcher.Fetch(path)
}

type BrokenReader struct{}

func (b BrokenReader) Read([]byte) (int, error) {
	return 0, errors.New("read on broken reader")
}

type BrokenFetcher struct{}

func (b BrokenFetcher) Fetch(path string) (io.ReadCloser, error) {
	return nil, errors.New("fetch on broken fetcher")
}

func NewFetchRecorder(fetcher Fetcher) FetchRecorder {
	return FetchRecorder{fetcher, []string{}}
}
