package util

import (
	"regexp"
	"strings"
)

func Format(s string) string {
	n := strings.TrimSpace(
		strings.ToLower(
			strings.ReplaceAll(s, " ", "_")))

	nonAlNum, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		panic(err)
	}

	consecutiveUnder, err := regexp.Compile("_{2,}")
	if err != nil {
		panic(err)
	}

	n = nonAlNum.ReplaceAllString(n, "")
	n = consecutiveUnder.ReplaceAllString(n, "_")
	n = strings.Trim(n, "_")
	return n
}
