package math_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/kamilsk/platform/pkg/math"
	"github.com/stretchr/testify/assert"
)

func TestReducer_Average(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		expected float64
	}{
		{"nil, zero average", nil, 0},
		{"empty, zero average", []int{}, 0},
		{"normal case", []int{1, 2, 3}, 2},
		{"fractional", []int{1, 2, 3, 4, 5, 6}, 3.5},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Reduce(tc.sequence...).Average())
		})
	}
}

func TestReducer_Median(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		expected float64
	}{
		{"nil, zero median", nil, 0},
		{"empty, zero median", nil, 0},
		{"even, sorted", []int{1, 2, 3, 4}, 2.5},
		{"even, unsorted", []int{2, 1, 4, 3}, 2.5},
		{"odd, sorted", []int{1, 2, 3, 4, 5}, 3},
		{"odd, unsorted", []int{3, 1, 2, 4, 5}, 3},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Reduce(tc.sequence...).Median())
		})
	}
}

func TestReducer_Sum(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		expected int
	}{
		{"nil, zero sum", nil, 0},
		{"empty, zero sum", []int{}, 0},
		{"positive sum", []int{1, 2, 3}, 6},
		{"negative sum", []int{-1, -2, -3}, -6},
		{"mixed, zero sum", []int{-1, -2, 3}, 0},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Reduce(tc.sequence...).Sum())
		})
	}
}

func TestSequence(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"constant", 5},
		{"random", rand.New(rand.NewSource(time.Now().UnixNano())).Int()},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Len(t, Sequence(tc.size), tc.size)
		})
	}
}
