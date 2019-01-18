package math

import (
	"math"
	"sort"
)

// Reduce wraps sequence to perform some aggregate operations above it.
func Reduce(sequence ...int) interface {
	// Average returns the average value of the sequence.
	Average() float64
	// Median returns the median value of the sequence.
	Median() float64
	// Sum returns the sum of the sequence.
	Sum() int
} {
	return reducer(sequence)
}

type reducer []int

// Average returns the average value of the sequence.
func (sequence reducer) Average() float64 {
	if len(sequence) == 0 {
		return 0
	}
	return float64(sequence.Sum()) / float64(len(sequence))
}

func (sequence reducer) Count() int {
	return len(sequence)
}

func (sequence reducer) Maximum() int {
	if len(sequence) == 0 {
		return 0
	}
	max := math.MinInt64
	for _, num := range sequence {
		if num > max {
			max = num
		}
	}
	return max
}

// Median returns the median value of the sequence.
func (sequence reducer) Median() float64 {
	size := len(sequence)
	if size == 0 {
		return 0
	}
	sorted := make([]int, size)
	copy(sorted, sequence)
	sort.Ints(sorted)
	if size%2 == 0 {
		return (float64(sorted[size/2-1]) + float64(sorted[size/2])) / 2
	}
	return float64(sorted[size/2])
}

func (sequence reducer) Minimum() int {
	if len(sequence) == 0 {
		return 0
	}
	min := math.MaxInt64
	for _, num := range sequence {
		if num < min {
			min = num
		}
	}
	return min
}

// Sum returns the sum of the sequence.
func (sequence reducer) Sum() int {
	sum := 0
	for _, num := range sequence {
		sum += num
	}
	return sum
}

// Sequence returns an empty slice with the specified size.
//
//     for range Sequence(5) {
//             // do something five times
//     }
//
func Sequence(size int) []struct{} {
	return make([]struct{}, size)
}
