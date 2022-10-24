package bconv

import (
	"reflect"
	"unsafe"
)

// unsafe convert Bytes to String.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(
		&reflect.StringHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
			Len:  len(b),
		}))
}

// unsafe convert String to Bytes.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
			Len:  len(s),
			Cap:  len(s),
		}))
}
