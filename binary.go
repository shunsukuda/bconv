package bconv

import (
	"reflect"
	"unsafe"
)

type Binary []byte
type String string
type U16 uint16
type U32 uint32
type U64 uint64
type I16 int16
type I32 int32
type I64 int64
type F32 float32
type F64 float64
type U16s []uint16
type U32s []uint32
type U64s []uint64
type I16s []int16
type I32s []int32
type I64s []int64
type F32s []float32
type F64s []float64

// BinToU16 convert []byte to uint16.
// ByteOrder are dependent on envionments.
func BinToU16s(b Binary) []uint16 {
	if len(b) >= 2 {
		return *(*[]uint16)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 2),
				Cap:  int(cap(b) / 2),
			}))
	} else {
		return nil
	}
}

// BinToU32 convert []byte to uint32.
// ByteOrder are dependent on envionments.
func BinToU32s(b Binary) []uint32 {
	if len(b) >= 4 {
		return *(*[]uint32)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 4),
				Cap:  int(cap(b) / 4),
			}))
	} else {
		return nil
	}
}

// BinToU64 convert []byte to uint64.
// ByteOrder are dependent on envionments.
func BinToU64s(b Binary) []uint64 {
	if len(b) >= 8 {
		return *(*[]uint64)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 8),
				Cap:  int(cap(b) / 8),
			}))
	} else {
		return nil
	}
}

// BinToI16 convert []byte to int16.
// ByteOrder are dependent on envionments.
func BinToI16s(b Binary) []int16 {
	if len(b) >= 2 {
		return *(*[]int16)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 2),
				Cap:  int(cap(b) / 2),
			}))
	} else {
		return nil
	}
}

// BinToI32 convert []byte to int32.
// ByteOrder are dependent on envionments.
func BinToI32s(b Binary) []int32 {
	if len(b) >= 4 {
		return *(*[]int32)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 4),
				Cap:  int(cap(b) / 4),
			}))
	} else {
		return nil
	}
}

// BinToI64 convert []byte to int64.
// ByteOrder are dependent on envionments.
func BinToI64s(b Binary) []int64 {
	if len(b) >= 8 {
		return *(*[]int64)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 8),
				Cap:  int(cap(b) / 8),
			}))
	} else {
		return nil
	}
}

// BinToF32 convert []byte to float32.
// ByteOrder are dependent on envionments.
func BinToF32s(b Binary) []float32 {
	if len(b) >= 4 {
		return *(*[]float32)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 4),
				Cap:  int(cap(b) / 4),
			}))
	} else {
		return nil
	}
}

// BinToF64 convert []byte to float64.
// ByteOrder are dependent on envionments.
func BinToF64s(b Binary) []float64 {
	if len(b) >= 8 {
		return *(*[]float64)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
				Len:  int(len(b) / 8),
				Cap:  int(cap(b) / 8),
			}))
	} else {
		return nil
	}
}

// U16sToBin convert []uint16 to []byte.
// ByteOrder are dependent on envionments.
func U16sToBin(v []uint16) Binary {
	if v != nil {
		return *(*Binary)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
				Len:  int(len(v) * 2),
				Cap:  int(cap(v) * 2),
			}))
	} else {
		return nil
	}
}

// U32sToBin convert []uint32 to []byte.
// ByteOrder are dependent on envionments.
func U32sToBin(v []uint32) Binary {
	if v != nil {
		return *(*Binary)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
				Len:  int(len(v) * 4),
				Cap:  int(cap(v) * 4),
			}))
	} else {
		return nil
	}
}

// U64sToBin convert []uint64 to []byte.
// ByteOrder are dependent on envionments.
func U64sToBin(v []uint64) Binary {
	if v != nil {
		return *(*Binary)(unsafe.Pointer(
			&reflect.SliceHeader{
				Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
				Len:  int(len(v) * 8),
				Cap:  int(cap(v) * 8),
			}))
	} else {
		return nil
	}
}
