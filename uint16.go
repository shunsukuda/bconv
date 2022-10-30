package bconv

import (
	"reflect"
	"unsafe"
)

// U16ToI16 convert uint16 to int16.
func U16ToI16(v uint16) int16 {
	return *(*int16)(unsafe.Pointer(&v))
}

// U16sToI16s convert []uint16 to []int16.
func U16sToI16s(v []uint16) []int16 {
	return *(*[]int16)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// I16ToU16 convert int16 to uint16.
func I16ToU16(v int16) uint16 {
	return *(*uint16)(unsafe.Pointer(&v))
}

// I16sToU16s convert []int16 to []uint16.
func I16sToU16s(v []int16) []uint16 {
	return *(*[]uint16)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
