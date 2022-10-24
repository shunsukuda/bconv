package bconv

import (
	"reflect"
	"unsafe"
)

// Uint64ToInt64 convert uint64 to int64.
func Uint64ToInt64(v uint64) int64 {
	return *(*int64)(unsafe.Pointer(&v))
}

// Uint64SliceToInt64Slice convert []uint64 to []int64.
func Uint64SliceToInt64Slice(v []uint64) []int64 {
	return *(*[]int64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Uint64ToFloat64 convert uint64 to float64.
func Uint64ToFloat64(v uint64) float64 {
	return *(*float64)(unsafe.Pointer(&v))
}

// Uint64SliceToFloat64Slice convert []uint64 to []float64.
func Uint64SliceToFloat64Slice(v []uint64) []float64 {
	return *(*[]float64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Int64ToUint64 convert int64 to uint64.
func Int64ToUint64(v int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

// Int64SliceToUint64Slice convert []int64 to []uint64.
func Int64SliceToUint64Slice(v []int64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Float64ToUint64 convert float64 to uint64.
func Float64ToUint64(v float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

// Float64SliceToUint64Slice convert []float64 to []uint64.
func Float64SliceToUint64Slice(v []float64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
