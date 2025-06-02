package helpers

import (
	"errors"
	"time"
)

const layoutDate = "2006-01-02"

func GetDayName(dateStr string) (string, error) {
	parsedDate, err := time.Parse(layoutDate, dateStr)
	if err != nil {
		return "", err
	}
	return parsedDate.Weekday().String(), nil
}

func ValidateAndNormalizeDateString(input string) (string, error) {
	if input == "" {
		return "", errors.New("date input is empty")
	}

	parsedTime, err := time.Parse(layoutDate, input)
	if err != nil {
		return "", err
	}

	return parsedTime.Format(layoutDate), nil
}
