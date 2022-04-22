package query_test

import (
	"testing"

	"github.com/olzemal/lsemoji/pkg/cache"
	"github.com/olzemal/lsemoji/pkg/query"
)

func TestSearch(t *testing.T) {
	c1 := cache.Cache{
		"smile": "😊",
	}
	c2 := cache.Cache{
		"smile":     "😊",
		"cat_smile": "😼",
	}
	tests := []struct {
		c cache.Cache
		q string
		r map[string]string
	}{
		{c: c1, q: "smile", r: map[string]string{"smile": "😊"}},
		{c: c2, q: "smile", r: map[string]string{"smile": "😊", "cat_smile": "😼"}},
	}

	for _, test := range tests {
		r := query.FilterContains(test.c, test.q)
		if !EqualMaps[string](r, test.r) {
			t.Errorf("Got: `%v`, Want; `%v`", r, test.r)
		}
	}
}

func EqualMaps[T comparable](a, b map[T]T) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
