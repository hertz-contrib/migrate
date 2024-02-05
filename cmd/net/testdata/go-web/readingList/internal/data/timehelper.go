package data

import "time"

func toTime(value string) (*time.Time, error) {
	timeValue, err := time.Parse("2006-01-02T15:04:05Z", value)
	if err != nil {
		return nil, err
	}
	return &timeValue, nil
}
