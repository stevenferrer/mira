package mira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtrs(t *testing.T) {
	s := ""
	assert.Equal(t, &s, StrPtr(s))

	b := true
	assert.Equal(t, &b, BoolPtr(b))

	i := int(0)
	assert.Equal(t, &i, IntPtr(i))

	i8 := int8(0)
	assert.Equal(t, &i8, Int8Ptr(i8))

	i16 := int16(0)
	assert.Equal(t, &i16, Int16Ptr(i16))

	i32 := int32(0)
	assert.Equal(t, &i32, Int32Ptr(i32))

	i64 := int64(0)
	assert.Equal(t, &i64, Int64Ptr(0))

	ui := uint(0)
	assert.Equal(t, &ui, UintPtr(ui))

	ui8 := uint8(0)
	assert.Equal(t, &ui8, Uint8Ptr(ui8))

	ui16 := uint16(0)
	assert.Equal(t, &ui16, Uint16Ptr(ui16))

	ui32 := uint32(0)
	assert.Equal(t, &ui32, Uint32Ptr(0))

	ui64 := uint64(0)
	assert.Equal(t, &ui64, Uint64Ptr(ui64))

	f32 := float32(0)
	assert.Equal(t, &f32, Float32Ptr(f32))

	f64 := float64(0)
	assert.Equal(t, &f64, Float64Ptr(f64))

	c64 := complex64(0)
	assert.Equal(t, &c64, Complex64Ptr(c64))

	c128 := complex128(0)
	assert.Equal(t, &c128, Complex128Ptr(c128))
}
