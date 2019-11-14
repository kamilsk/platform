// Package pointer provides helpers to get pointers to values of built-in primitives.
//
// Based on https://github.com/AlekSi/pointer.
package pointer

import "time"

// ToBool returns pointer to boolean primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToBool.
func ToBool(b bool) *bool { return &b }

// ToByte returns pointer to byte primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToByte.
func ToByte(b byte) *byte { return &b }

// ToFloat32 returns pointer to float32 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToFloat32.
func ToFloat32(f float32) *float32 { return &f }

// ToFloat64 returns pointer to float64 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToFloat64.
func ToFloat64(f float64) *float64 { return &f }

// ToInt returns pointer to int primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToInt.
func ToInt(i int) *int { return &i }

// ToInt16 returns pointer to int16 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToInt16.
func ToInt16(i int16) *int16 { return &i }

// ToInt32 returns pointer to int32 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToInt32.
func ToInt32(i int32) *int32 { return &i }

// ToInt64 returns pointer to int64 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToInt64.
func ToInt64(i int64) *int64 { return &i }

// ToInt8 returns pointer to int8 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToInt8.
func ToInt8(i int8) *int8 { return &i }

// ToRune returns pointer to rune primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToRune.
func ToRune(r rune) *rune { return &r }

// ToString returns pointer to string primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToString.
func ToString(s string) *string { return &s }

// ToTime returns pointer to time.Time.
// Deprecated: use original package or go.octolab.org/pointer.ToTime.
func ToTime(t time.Time) *time.Time { return &t }

// ToUint returns pointer to uint primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUint.
func ToUint(u uint) *uint { return &u }

// ToUint16 returns pointer to uint16 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUint16.
func ToUint16(u uint16) *uint16 { return &u }

// ToUint32 returns pointer to uint32 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUint32.
func ToUint32(u uint32) *uint32 { return &u }

// ToUint64 returns pointer to uint64 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUint64.
func ToUint64(u uint64) *uint64 { return &u }

// ToUint8 returns pointer to uint8 primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUint8.
func ToUint8(u uint8) *uint8 { return &u }

// ToUintptr returns pointer to uintptr primitive.
// Deprecated: use original package or go.octolab.org/pointer.ToUintptr.
func ToUintptr(u uintptr) *uintptr { return &u }
