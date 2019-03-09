package genome

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
