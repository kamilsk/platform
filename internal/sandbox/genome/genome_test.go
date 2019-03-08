package genome_test

import (
	"testing"

	. "github.com/kamilsk/platform/internal/sandbox/genome"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name   string
		origin []T
	}{
		{
			"nil",
			nil,
		},
		{
			"empty",
			[]T{},
		},
		{
			"not empty",
			[]T{1, 2, 3},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.origin, Copy(tc.origin))
		})
	}
}

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		origin   []T
		from, to int
		expected []T
	}{
		{
			"left",
			[]T{1, 2, 3, 4, 5},
			0, 2,
			[]T{3, 4, 5},
		},
		{
			"right",
			[]T{1, 2, 3, 4, 5},
			3, 5,
			[]T{1, 2, 3},
		},
		{
			"center",
			[]T{1, 2, 3, 4, 5},
			1, 4,
			[]T{1, 5},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Cut(tc.origin, tc.from, tc.to))
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		src      []T
		position int
		expected []T
	}{
		{
			"first",
			[]T{1, 2, 3},
			0,
			[]T{2, 3},
		},
		{
			"center",
			[]T{1, 2, 3},
			1,
			[]T{1, 3},
		},
		{
			"last",
			[]T{1, 2, 3},
			2,
			[]T{1, 2},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Delete(tc.src, tc.position))
		})
	}
}
