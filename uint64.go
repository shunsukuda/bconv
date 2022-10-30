package bconv

import (
	"reflect"
	"unsafe"
)

// U64ToI64 convert uint64 to int64.
func U64ToI64(v uint64) int64 {
	return *(*int64)(unsafe.Pointer(&v))
}

// U64sToI64s convert []uint64 to []int64.
func U64sToI64s(v []uint64) []int64 {
	return *(*[]int64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// U64ToF64 convert uint64 to float64.
func U64ToF64(v uint64) float64 {
	return *(*float64)(unsafe.Pointer(&v))
}

// U64sToF64s convert []uint64 to []float64.
func U64sToF64s(v []uint64) []float64 {
	return *(*[]float64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// I64ToU64 convert int64 to uint64.
func I64ToU64(v int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

// I64sToU64s convert []int64 to []uint64.
func I64sToU64s(v []int64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// F64ToU64 convert float64 to uint64.
func F64ToU64(v float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

// F64sToU64s convert []float64 to []uint64.
func F64sToU64s(v []float64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
