package helpers

import "time"

func GetDayName(dateStr string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	return parsedDate.Weekday().String(), nil
}
