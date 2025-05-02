package helpers

import (
	"errors"
	"time"
)

func ParseBirthdate(dateStr string) (*time.Time, error) {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return nil, errors.New("invalid date format, expected dd-mm-yyyy")
	}
	return &t, nil
}
