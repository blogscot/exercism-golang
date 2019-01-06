package gigasecond

import (
	"time"
)

const (
	gigaSecond = 1000000000 * time.Second
)

// AddGigasecond adds 10e9 seconds to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
