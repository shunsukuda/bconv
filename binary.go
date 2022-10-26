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

// BinToU16 convert []byte(little-endian) to uint16.
func BinLEToU16(b Binary) uint16 {
	if len(b) >= 2 {
		return uint16(b[1])<<8 | uint16(b[0])
	} else {
		return 0
	}
}

// BinToU16 convert []byte(big-endian) to uint16.
func BinBEToU16(b Binary) uint16 {
	if len(b) >= 2 {
		return uint16(b[1]) | uint16(b[0])<<8
	} else {
		return 0
	}
}

// BinToU32 convert []byte(little-endian) to uint32.
func BinLEToU32(b Binary) uint32 {
	if len(b) >= 4 {
		return uint32(b[3])<<24 |
			uint32(b[2])<<16 |
			uint32(b[1])<<8 |
			uint32(b[0])
	} else {
		return 0
	}
}

// BinToU32 convert []byte(big-endian) to uint32.
func BinBEToU32(b Binary) uint32 {
	if len(b) >= 4 {
		return uint32(b[3]) |
			uint32(b[2])<<8 |
			uint32(b[1])<<16 |
			uint32(b[0])<<24
	} else {
		return 0
	}
}

// BinToU64 convert []byte(little-endian) to uint64.
func BinLEToU64(b Binary) uint64 {
	if len(b) >= 8 {
		return uint64(b[7])<<56 |
			uint64(b[6])<<48 |
			uint64(b[5])<<40 |
			uint64(b[4])<<32 |
			uint64(b[3])<<24 |
			uint64(b[2])<<16 |
			uint64(b[1])<<8 |
			uint64(b[0])
	} else {
		return 0
	}
}

// BinToU64 convert []byte(big-endian) to uint64.
func BinBEToU64(b Binary) uint64 {
	if len(b) >= 8 {
		_ = b[7]
		return uint64(b[7]) |
			uint64(b[6])<<8 |
			uint64(b[5])<<16 |
			uint64(b[4])<<24 |
			uint64(b[3])<<32 |
			uint64(b[2])<<40 |
			uint64(b[1])<<48 |
			uint64(b[0])<<56
	} else {
		return 0
	}
}

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

// U16ToBinLE convert uint16 to []byte(little-endian).
func U16ToBinLE(v uint16) Binary {
	return Binary{
		byte(v),
		byte(v >> 8),
	}
}

// U16ToBinBE convert uint16 to []byte(big-endian).
func U16ToBinBE(v uint16) Binary {
	return Binary{
		byte(v >> 8),
		byte(v),
	}
}

// U32ToBinLE convert uint32 to []byte(little-endian).
func U32ToBinLE(v uint32) Binary {
	return Binary{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
}

// U32ToBinBE convert uint32 to []byte(big-endian).
func U32ToBinBE(v uint32) Binary {
	return Binary{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

// U64ToBinLE convert uint64 to []byte(little-endian).
func U64ToBinLE(v uint64) Binary {
	return Binary{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}

// U64ToBinBE convert uint64 to []byte(big-endian).
func U64ToBinBE(v uint64) Binary {
	return Binary{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}
