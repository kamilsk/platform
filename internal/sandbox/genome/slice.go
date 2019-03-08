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
	return append(src[:at], append(make([]T, size), src[at:]...)...)
}
