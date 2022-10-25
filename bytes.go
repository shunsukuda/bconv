package bconv

import (
	"reflect"
	"unsafe"
)

// BytesToUint16 convert []byte to uint16(ignore endian)
func BytesToUint16(b []byte) uint16 {
	if len(b) >= 2 {
		return *(*uint16)(unsafe.Pointer(
			reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  len(b),
				Cap:  len(b),
			}.Data))
	} else {
		return 0
	}
}

// BytesToUint32 convert []byte to uint32(ignore endian)
func BytesToUint32(b []byte) uint32 {
	if len(b) >= 4 {
		return *(*uint32)(unsafe.Pointer(
			reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  len(b),
				Cap:  len(b),
			}.Data))
	} else {
		return 0
	}
}

// BytesToUint64 convert []byte to uint64(ignore endian)
func BytesToUint64(b []byte) uint64 {
	if len(b) >= 8 {
		return *(*uint64)(unsafe.Pointer(
			reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  len(b),
				Cap:  len(b),
			}.Data))
	} else {
		return 0
	}
}
