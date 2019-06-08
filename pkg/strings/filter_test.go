package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/strings"
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

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		expected []string
	}{
		{name: "nothing to pass"},
		{"simple usage", []string{"one", "", "two", "", "three"}, []string{"one", "two", "three"}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NotEmpty(tc.strings))
		})
	}
}
