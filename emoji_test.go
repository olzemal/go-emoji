package emoji_test

import (
	"testing"

	emoji "github.com/olzemal/lsemoji"
)

func TestGenerateCache(t *testing.T) {
	c, _ := emoji.GenerateCache()
	emoji.ExportCache("", c)
}

func TestFromString(t *testing.T) {

}

func TestBestMatch(t *testing.T) {

}
