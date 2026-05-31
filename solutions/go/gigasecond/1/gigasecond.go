package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	const gigaSecond = 1e9 * time.Second
	return t.Add(gigaSecond)
}
