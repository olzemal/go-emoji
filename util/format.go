package util

import (
	"regexp"
	"strings"
)

func Format(s string) string {
	n := strings.TrimSpace(
		strings.ToLower(
			strings.ReplaceAll(s, " ", "_")))

	reg, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		panic(err)
	}
	n = reg.ReplaceAllString(n, "")
	n = strings.Trim(n, "_")
	return n
}
