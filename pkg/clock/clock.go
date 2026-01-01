package clock

import (
	"errors"
	"strconv"
	"time"
)

var ErrBadTime = errors.New("bad time string")

// IsDay returns true if t is considered daytime.
func IsDay(t time.Time) bool {
	h := t.Hour()
	return h >= 7 && h < 20
}

// Parse parses a unix timestamp into a time.Time
func Parse(ts string) (time.Time, error) {
	var t time.Time
	if ts == "" {
		return t, ErrBadTime
	}
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return t, err
	}
	t = time.Unix(i, 0)
	return t, nil
}
