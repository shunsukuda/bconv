package bconv

import (
	"reflect"
	"unsafe"
)

// U8ToI8 convert uint8 to int8.
func U8ToI8(v uint8) int8 {
	return *(*int8)(unsafe.Pointer(&v))
}

// U8sToI8s convert []uint8 to []int8.
func U8sToI8s(v []uint8) []int8 {
	return *(*[]int8)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// I8ToU8 convert int8 to uint8.
func I8ToU8(v int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&v))
}

// I8sToU8se convert []int8 to []uint8.
func I8sToU8s(v []int8) []uint8 {
	return *(*[]uint8)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
