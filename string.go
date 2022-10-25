package bconv

import (
	"reflect"
	"unsafe"
)

type Binary []byte

// BinToStr convert Binary to String.
func BinToStr(b Binary) string {
	return *(*string)(unsafe.Pointer(
		&reflect.StringHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
			Len:  len(b),
		}))
}

// SafeBinToStr safe convert Binary to String.
func SafeBinToStr(b Binary) string {
	b2 := make(Binary, len(b))
	copy(b2, b)
	return BinToStr(b2)
}

// StrToBin convert String to Binary.
func StrToBin(s string) Binary {
	return *(*Binary)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
			Len:  len(s),
			Cap:  len(s),
		}))
}

// SafeStrToBin safe convert String to Binary.
func SafeStrToBin(s string) Binary {
	b := StrToBin(s)
	b2 := make(Binary, len(b))
	copy(b2, b)
	return b2
}
