package main

import "io"

// Fetcher encapsulates the act of opening a named resource
// (such as a url or file path) and getting a reader for that resource.
type Fetcher interface {
	Fetch(path string) (io.ReadCloser, error)
}
