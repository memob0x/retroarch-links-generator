package utils

import "regexp"

var regexWinOSFilenames *regexp.Regexp = regexp.MustCompile(`\\|\/|:|\*|\?|"|<|>|\|`)

func GetValidWinOsFilename(value string) string {
	return regexWinOSFilenames.ReplaceAllString(value, "-")
}
