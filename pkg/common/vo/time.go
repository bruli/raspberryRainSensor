package vo

import (
	"errors"
	"strconv"
	"time"
)

var ErrInvalidZeroTime = errors.New("invalid zero time to build Time object")

type Time time.Time

func ParseFromTime(t time.Time) (Time, error) {
	if t.IsZero() {
		return Time{}, ErrInvalidZeroTime
	}
	return Time(t), nil
}

func (t Time) EpochString() string {
	unix := time.Time(t).Unix()
	return strconv.Itoa(int(unix))
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func TimeNow() Time {
	return Time(time.Now())
}
