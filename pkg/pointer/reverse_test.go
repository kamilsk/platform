package pointer_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/pointer"
)

func TestValueOfBool(t *testing.T) {
	tests := []struct {
		b *bool
		v bool
	}{
		{},
		{ToBool(true), true},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfBool(test.b))
	}
}

func TestValueOfByte(t *testing.T) {
	tests := []struct {
		b *byte
		v byte
	}{
		{},
		{ToByte(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfByte(test.b))
	}
}

func TestValueOfFloat32(t *testing.T) {
	tests := []struct {
		f *float32
		v float32
	}{
		{},
		{ToFloat32(7.7), 7.7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfFloat32(test.f))
	}
}

func TestValueOfFloat64(t *testing.T) {
	tests := []struct {
		f *float64
		v float64
	}{
		{},
		{ToFloat64(7.7), 7.7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfFloat64(test.f))
	}
}

func TestValueOfInt(t *testing.T) {
	tests := []struct {
		i *int
		v int
	}{
		{},
		{ToInt(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfInt(test.i))
	}
}

func TestValueOfInt16(t *testing.T) {
	tests := []struct {
		i *int16
		v int16
	}{
		{},
		{ToInt16(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfInt16(test.i))
	}
}

func TestValueOfInt32(t *testing.T) {
	tests := []struct {
		i *int32
		v int32
	}{
		{},
		{ToInt32(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfInt32(test.i))
	}
}

func TestValueOfInt64(t *testing.T) {
	tests := []struct {
		i *int64
		v int64
	}{
		{},
		{ToInt64(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfInt64(test.i))
	}
}

func TestValueOfInt8(t *testing.T) {
	tests := []struct {
		i *int8
		v int8
	}{
		{},
		{ToInt8(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfInt8(test.i))
	}
}

func TestValueOfRune(t *testing.T) {
	tests := []struct {
		r *rune
		v rune
	}{
		{},
		{ToRune('a'), 'a'},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfRune(test.r))
	}
}

func TestValueOfString(t *testing.T) {
	tests := []struct {
		s *string
		v string
	}{
		{},
		{ToString("a"), "a"},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfString(test.s))
	}
}

func TestValueOfTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		t *time.Time
		v time.Time
	}{
		{},
		{ToTime(now), now},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfTime(test.t))
	}
}

func TestValueOfUint(t *testing.T) {
	tests := []struct {
		u *uint
		v uint
	}{
		{},
		{ToUint(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUint(test.u))
	}
}

func TestValueOfUint16(t *testing.T) {
	tests := []struct {
		u *uint16
		v uint16
	}{
		{},
		{ToUint16(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUint16(test.u))
	}
}

func TestValueOfUint32(t *testing.T) {
	tests := []struct {
		u *uint32
		v uint32
	}{
		{},
		{ToUint32(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUint32(test.u))
	}
}

func TestValueOfUint64(t *testing.T) {
	tests := []struct {
		u *uint64
		v uint64
	}{
		{},
		{ToUint64(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUint64(test.u))
	}
}

func TestValueOfUint8(t *testing.T) {
	tests := []struct {
		u *uint8
		v uint8
	}{
		{},
		{ToUint8(7), 7},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUint8(test.u))
	}
}

func TestValueOfUintptr(t *testing.T) {
	tests := []struct {
		u *uintptr
		v uintptr
	}{
		{},
		{ToUintptr(0x1), 0x1},
	}
	for _, test := range tests {
		assert.Equal(t, test.v, ValueOfUintptr(test.u))
	}
}
