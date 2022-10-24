package bconv

import (
	"reflect"
	"unsafe"
)

// Uint8ToInt8 convert uint8 to int8.
func Uint8ToInt8(v uint8) int8 {
	return *(*int8)(unsafe.Pointer(&v))
}

// Uint8SliceToInt8Slice convert []uint8 to []int8.
func Uint8SliceToInt8Slice(v []uint8) []int8 {
	return *(*[]int8)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Int8ToUint8 convert int8 to uint8.
func Int8ToUint8(v int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&v))
}

// Int8SliceToUint8Slice convert []int8 to []uint8.
func Int8SliceToUint8Slice(v []int8) []uint8 {
	return *(*[]uint8)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
