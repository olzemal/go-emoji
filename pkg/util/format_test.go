package util_test

import (
	"testing"

	"github.com/olzemal/go-emoji/pkg/util"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		I string
		O string
	}{
		{I: "hello", O: "hello"},
		{I: "hello world", O: "hello_world"},
		{I: "hello & world", O: "hello_world"},
		{I: "a (world)", O: "a_world"},
		{I: "hello     world", O: "hello_world"},
		{I: "Hello World", O: "hello_world"},
	}

	for _, test := range tests {
		got := util.Format(test.I)
		if got != test.O {
			t.Errorf("Got: %v, Want: %v", got, test.O)
		}
	}
}
