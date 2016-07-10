package main

import "time"

// PrettyDuration turns a duration into a "pretty" human-readable string
func PrettyDuration(duration time.Duration) string {
	// Round to seconds
	return (time.Second * (duration / time.Second)).String()
}
