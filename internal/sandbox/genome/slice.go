package genome

type T int64

func Copy(origin []T) (replica []T) {
	if origin == nil {
		return nil
	}
	replica = make([]T, len(origin))
	copy(replica, origin)
	return replica
}

func Cut(src []T, from, to int) []T {
	return append(src[:from], src[to:]...)
}
