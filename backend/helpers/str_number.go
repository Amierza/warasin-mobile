package helpers

import "regexp"

func IsValidSTRNumber(str string) bool {
	pattern := `^ST\d{8}$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
