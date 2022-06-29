package utils

import (
	"regexp"
)

func SplitStringWithRegex(value string, pattern string) []string {
	var regExp *regexp.Regexp = regexp.MustCompile(pattern)

	return regExp.Split(value, -1)
}
