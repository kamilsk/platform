// Package pointer provides helpers to get pointers to values of built-in primitives.
//
// Based on https://github.com/AlekSi/pointer.
package pointer

// ToBool returns pointer to boolean primitive.
func ToBool(b bool) *bool { return &b }

// ToByte returns pointer to byte primitive.
func ToByte(b byte) *byte { return &b }

// ToFloat32 returns pointer to float32 primitive.
func ToFloat32(f float32) *float32 { return &f }

// ToFloat64 returns pointer to float64 primitive.
func ToFloat64(f float64) *float64 { return &f }

// ToInt returns pointer to int primitive.
func ToInt(i int) *int { return &i }

// ToInt16 returns pointer to int16 primitive.
func ToInt16(i int16) *int16 { return &i }

// ToInt32 returns pointer to int32 primitive.
func ToInt32(i int32) *int32 { return &i }

// ToInt64 returns pointer to int64 primitive.
func ToInt64(i int64) *int64 { return &i }

// ToInt8 returns pointer to int8 primitive.
func ToInt8(i int8) *int8 { return &i }

// ToRune returns pointer to rune primitive.
func ToRune(r rune) *rune { return &r }

// ToString returns pointer to string primitive.
func ToString(s string) *string { return &s }

// ToUint returns pointer to uint primitive.
func ToUint(u uint) *uint { return &u }

// ToUint16 returns pointer to uint16 primitive.
func ToUint16(u uint16) *uint16 { return &u }

// ToUint32 returns pointer to uint32 primitive.
func ToUint32(u uint32) *uint32 { return &u }

// ToUint64 returns pointer to uint64 primitive.
func ToUint64(u uint64) *uint64 { return &u }

// ToUint8 returns pointer to uint8 primitive.
func ToUint8(u uint8) *uint8 { return &u }

// ToUintptr returns pointer to uintptr primitive.
func ToUintptr(u uintptr) *uintptr { return &u }
