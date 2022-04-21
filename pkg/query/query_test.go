package query_test

import (
	"fmt"
	"testing"

	"github.com/olzemal/lsemoji/pkg/cache"
	"github.com/olzemal/lsemoji/pkg/query"
)

func TestSearch(t *testing.T) {
	c1 := cache.Cache{
		"smile": "😊",
	}
	tests := []struct {
		c cache.Cache
		q string
		r []string
	}{
		{c: c1, q: "smile", r: []string{"😊"}},
	}

	for _, test := range tests {
		r := query.Search(test.q, test.c)
		fmt.Println(r)
		if !ExactMatch(r, test.r) {
			t.Errorf("Got: `%v`, Want; `%v`", r, test.r)
		}
	}
}

func ExactMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
