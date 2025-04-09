package timex

import (
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func ParseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		return lt, nil
	}
}
