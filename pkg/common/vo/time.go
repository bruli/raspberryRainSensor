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

func (t Time) HourStr() string {
	return time.Time(t).Format("15:04")
}

func (t Time) Add(d time.Duration) Time {
	dt := time.Time(t).Add(d)
	return Time(dt)
}

func (t Time) AddDate(years, months, days int) Time {
	dt := time.Time(t).AddDate(years, months, days)
	return Time(dt)
}

func TimeNow() Time {
	return Time(time.Now())
}
