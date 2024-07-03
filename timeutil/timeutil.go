package timeutil

import (
	"time"
)

func ParseRangeDate(start string, end string) (time.Time, time.Time, error) {
	startDate, err := time.ParseInLocation(time.DateOnly, start, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endDate, err := time.ParseInLocation(time.DateOnly, end, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startDate, endDate, nil
}

func ParseRangeTime(start string, end string) (time.Time, time.Time, error) {
	startTime, err := time.ParseInLocation(time.DateTime, start, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endTime, err := time.ParseInLocation(time.DateTime, end, time.Local)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startTime, endTime, nil
}
