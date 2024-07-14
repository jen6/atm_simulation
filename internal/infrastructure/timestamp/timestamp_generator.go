package timestamp

import "time"

type TimestampGenerator struct {
	fixedTimestamp *time.Time
}

func (t TimestampGenerator) Now() time.Time {
	if t.fixedTimestamp != nil {
		return *t.fixedTimestamp
	}
	return time.Now()
}
