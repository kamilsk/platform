package pointer

import "time"

// ValueOfBool returns value of pointer or its default value if pointer is nil.
func ValueOfBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// ValueOfByte returns value of pointer or its default value if pointer is nil.
func ValueOfByte(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

// ValueOfFloat32 returns value of pointer or its default value if pointer is nil.
func ValueOfFloat32(f *float32) float32 {
	if f == nil {
		return 0.0
	}
	return *f
}

// ValueOfFloat64 returns value of pointer or its default value if pointer is nil.
func ValueOfFloat64(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// ValueOfInt returns value of pointer or its default value if pointer is nil.
func ValueOfInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt16 returns value of pointer or its default value if pointer is nil.
func ValueOfInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt32 returns value of pointer or its default value if pointer is nil.
func ValueOfInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt64 returns value of pointer or its default value if pointer is nil.
func ValueOfInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt8 returns value of pointer or its default value if pointer is nil.
func ValueOfInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfRune returns value of pointer or its default value if pointer is nil.
func ValueOfRune(r *rune) rune {
	if r == nil {
		return 0x00
	}
	return *r
}

// ValueOfString returns value of pointer or its default value if pointer is nil.
func ValueOfString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ValueOfTime returns value of pointer or its default value if pointer is nil.
func ValueOfTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

// ValueOfUint returns value of pointer or its default value if pointer is nil.
func ValueOfUint(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint16 returns value of pointer or its default value if pointer is nil.
func ValueOfUint16(u *uint16) uint16 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint32 returns value of pointer or its default value if pointer is nil.
func ValueOfUint32(u *uint32) uint32 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint64 returns value of pointer or its default value if pointer is nil.
func ValueOfUint64(u *uint64) uint64 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint8 returns value of pointer or its default value if pointer is nil.
func ValueOfUint8(u *uint8) uint8 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUintptr returns value of pointer or its default value if pointer is nil.
func ValueOfUintptr(u *uintptr) uintptr {
	if u == nil {
		return 0x0
	}
	return *u
}
