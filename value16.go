package bconv

import (
	"reflect"
	"unsafe"
)

// Uint16ToInt16 convert uint16 to int16.
func Uint16ToInt16(v uint16) int16 {
	return *(*int16)(unsafe.Pointer(&v))
}

// Uint16SliceToInt16Slice convert []uint16 to []int16.
func Uint16SliceToInt16Slice(v []uint16) []int16 {
	return *(*[]int16)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Int16ToUint16 convert int16 to uint16.
func Int16ToUint16(v int16) uint16 {
	return *(*uint16)(unsafe.Pointer(&v))
}

// Int16SliceToUint16Slice convert []int16 to []uint16.
func Int16SliceToUint16Slice(v []int16) []uint16 {
	return *(*[]uint16)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
