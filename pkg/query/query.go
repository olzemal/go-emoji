package query

import (
	"regexp"

	"github.com/olzemal/lsemoji/pkg/cache"
)

func Search(query string, c cache.Cache) []string {
	var results []string
	substr, _ := regexp.Compile(".*" + query + ".*")
	for k := range c {
		if substr.MatchString(k) {
			results = append(results, c[k])
		}
	}
	return results
}
