package genome

import "math/rand"

func Copy(src []T) []T {
	replica := make([]T, len(src))
	copy(replica, src)
	return replica
}

func Cut(src []T, from, to int) []T {
	return append(src[:from], src[to:]...)
}

func Delete(src []T, i int) []T {
	return append(src[:i], src[i+1:]...)
}

func Expand(src []T, at, size int) []T {
	return append(src[:at], append(make([]T, size), src[at:]...)...)
}

func Extend(src []T, size int) []T {
	return append(src, make([]T, size)...)
}

func Insert(dst []T, t T, at int) []T {
	return append(dst[:at], append([]T{t}, dst[at:]...)...)
}

func InsertVector(dst []T, v []T, at int) []T {
	return append(dst[:at], append(v, dst[at:]...)...)
}

func Push(dst []T, t T) []T {
	return append(dst, t)
}

func Pop(dst []T) (T, []T) {
	last := len(dst) - 1
	return dst[last], dst[:last]
}

func Unshift(dst []T, t T) []T {
	return append([]T{t}, dst...)
}

func Shift(dst []T) (T, []T) {
	return dst[0], dst[1:]
}

func Filter(src []T, filter func(int, T) bool) []T {
	var def T
	filtered := src[:0]
	for i, t := range src {
		if filter(i, t) {
			filtered = append(filtered, t)
		}
	}
	for i, total := len(filtered), len(src); i < total; i++ {
		src[i] = def
	}
	return filtered
}

func FilterByKey(src []T, filter func(int) bool) []T {
	return Filter(src, func(i int, _ T) bool { return filter(i) })
}

func FilterByValue(src []T, filter func(T) bool) []T {
	return Filter(src, func(_ int, t T) bool { return filter(t) })
}

func Reverse(src []T) {
	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
}

func Shuffle(src []T, rand *rand.Rand) {
	for i := len(src) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		src[i], src[j] = src[j], src[i]
	}
}
