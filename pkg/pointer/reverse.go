package pointer

import "time"

// ValueOfBool returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetBool or go.octolab.org/pointer.ValueOfBool.
func ValueOfBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// ValueOfByte returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetByte or go.octolab.org/pointer.ValueOfByte.
func ValueOfByte(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

// ValueOfFloat32 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetFloat32 or go.octolab.org/pointer.ValueOfFloat32.
func ValueOfFloat32(f *float32) float32 {
	if f == nil {
		return 0.0
	}
	return *f
}

// ValueOfFloat64 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetFloat64 or go.octolab.org/pointer.ValueOfFloat64.
func ValueOfFloat64(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// ValueOfInt returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetInt or go.octolab.org/pointer.ValueOfInt.
func ValueOfInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt16 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetInt16 or go.octolab.org/pointer.ValueOfInt16.
func ValueOfInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt32 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetInt32 or go.octolab.org/pointer.ValueOfInt32.
func ValueOfInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt64 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetInt64 or go.octolab.org/pointer.ValueOfInt64.
func ValueOfInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt8 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetInt8 or go.octolab.org/pointer.ValueOfInt8.
func ValueOfInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfRune returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetRune or go.octolab.org/pointer.ValueOfRune.
func ValueOfRune(r *rune) rune {
	if r == nil {
		return 0x00
	}
	return *r
}

// ValueOfString returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetString or go.octolab.org/pointer.ValueOfString.
func ValueOfString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ValueOfTime returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetTime or go.octolab.org/pointer.ValueOfTime.
func ValueOfTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

// ValueOfUint returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUint or go.octolab.org/pointer.ValueOfUint.
func ValueOfUint(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint16 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUint16 or go.octolab.org/pointer.ValueOfUint16.
func ValueOfUint16(u *uint16) uint16 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint32 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUint32 or go.octolab.org/pointer.ValueOfUint32.
func ValueOfUint32(u *uint32) uint32 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint64 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUint64 or go.octolab.org/pointer.ValueOfUint64.
func ValueOfUint64(u *uint64) uint64 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint8 returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUint8 or go.octolab.org/pointer.ValueOfUint8.
func ValueOfUint8(u *uint8) uint8 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUintptr returns value of pointer or its default value if pointer is nil.
// Deprecated: use github.com/AlekSi/pointer.GetUintptr or go.octolab.org/pointer.ValueOfUintptr.
func ValueOfUintptr(u *uintptr) uintptr {
	if u == nil {
		return 0x0
	}
	return *u
}
