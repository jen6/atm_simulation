package application

import "time"

type TimestampGenerator interface {
	Now() time.Time
}
