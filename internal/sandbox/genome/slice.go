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
