package genome

type T int64

func Copy(src []T) []T {
	if src == nil {
		return nil
	}
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
	// optimizations:
	// - nothing to do if size is zero
	// - use Extend if at is the last element
	// - allocate + copy if at is the first element
	return append(src[:at], append(make([]T, size), src[at:]...)...)
}

func Extend(src []T, size int) []T {
	return append(src, make([]T, size)...)
}
