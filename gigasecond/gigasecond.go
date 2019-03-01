// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 10^9 seconds to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
