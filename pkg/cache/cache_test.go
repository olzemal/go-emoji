package cache_test

import (
	"testing"

	cache "github.com/olzemal/go-emoji/pkg/cache"
)

func TestGenerateCache(t *testing.T) {
	_, err := cache.Generate()
	if err != nil {
		t.Errorf("Failed to generate Cache: %v", err)
	}
}
