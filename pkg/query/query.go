package query

import (
	"strings"
)

func FilterContains(m map[string]string, s string) map[string]string {
	results := make(map[string]string)
	for k, v := range m {
		if strings.Contains(k, s) {
			results[k] = v
		}
	}
	return results
}
