package genome

type T int64

func Copy(origin []T) []T {
	if origin == nil {
		return nil
	}
	replica := make([]T, len(origin))
	copy(replica, origin)
	return replica
}

func Cut(src []T, from, to int) []T {
	return append(src[:from], src[to:]...)
}

func Delete(src []T, i int) []T {
	return append(src[:i], src[i+1:]...)
}
