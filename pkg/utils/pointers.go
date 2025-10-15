package utils

import (
	"time"
	"unsafe"
)

func PtrString(s string) *string {
	return &s
}

func PtrInt(i int) *int {
	return &i
}

func PtrUint(i uint) *uint {
	return &i
}

func PtrFloat64(f float64) *float64 {
	return &f
}

func PtrBool(b bool) *bool {
	return &b
}

func PtrTime(b time.Time) *time.Time {
	return &b
}

func PtrInterface(i interface{}) *interface{} {
	return &i
}
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Float32SliceToByteSlice(floats []float32) []byte {
	byteSlice := make([]byte, len(floats)*4)
	for i, f := range floats {
		bits := uint32FromFloat32(f)
		byteSlice[i*4] = byte(bits)
		byteSlice[i*4+1] = byte(bits >> 8)
		byteSlice[i*4+2] = byte(bits >> 16)
		byteSlice[i*4+3] = byte(bits >> 24)
	}
	return byteSlice
}

func uint32FromFloat32(f float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&f))
}
func Float64SliceToByteSlice(floats []float64) []byte {
	byteSlice := make([]byte, len(floats)*8)
	for i, f := range floats {
		bits := uint64FromFloat64(f)
		byteSlice[i*8] = byte(bits)
		byteSlice[i*8+1] = byte(bits >> 8)
		byteSlice[i*8+2] = byte(bits >> 16)
		byteSlice[i*8+3] = byte(bits >> 24)
		byteSlice[i*8+4] = byte(bits >> 32)
		byteSlice[i*8+5] = byte(bits >> 40)
		byteSlice[i*8+6] = byte(bits >> 48)
		byteSlice[i*8+7] = byte(bits >> 56)
	}
	return byteSlice
}

func uint64FromFloat64(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Uint16SliceToByteSlice(u16s []uint16) []byte {
	byteSlice := make([]byte, len(u16s)*2)
	for i, u := range u16s {
		byteSlice[i*2] = byte(u)
		byteSlice[i*2+1] = byte(u >> 8)
	}
	return byteSlice
}
