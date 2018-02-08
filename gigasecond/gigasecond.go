// Calculates the moment when someone has lived for 10^9 seconds.
package gigasecond

import "time"

const gigasecond = time.Second * 1000000000

// AddGigasecond adds a gigasecond to time t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
