package strings_test

import (
	"testing"

	. "github.com/kamilsk/platform/pkg/strings"
	"github.com/stretchr/testify/assert"
)

func TestFirstValid(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		expected string
	}{
		{name: "nothing to pass"},
		{"simple usage", []string{"", "", "third"}, "third"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, FirstValid(tc.strings...))
		})
	}
}
