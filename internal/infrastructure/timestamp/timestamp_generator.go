package timestamp

import "time"

type TimestampGenerator struct {
	FixedTimestamp *time.Time
}

func (t TimestampGenerator) Now() time.Time {
	if t.FixedTimestamp != nil {
		return *t.FixedTimestamp
	}
	return time.Now()
}
