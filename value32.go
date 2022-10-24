package bconv

import (
	"reflect"
	"unsafe"
)

// Uint32ToInt32 convert uint32 to int32.
func Uint32ToInt32(v uint32) int32 {
	return *(*int32)(unsafe.Pointer(&v))
}

// Uint32SliceToInt32Slice convert []uint32 to []int32.
func Uint32SliceToInt32Slice(v []uint32) []int32 {
	return *(*[]int32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Uint32ToFloat32 convert uint32 to float32.
func Uint32ToFloat32(v uint32) float32 {
	return *(*float32)(unsafe.Pointer(&v))
}

// Uint32SliceToFloat32Slice convert []uint32 to []float32.
func Uint32SliceToFloat32Slice(v []uint32) []float32 {
	return *(*[]float32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Int32ToUint32 convert int32 to uint32.
func Int32ToUint32(v int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}

// Int32SliceToUint32Slice convert []int32 to []uint32.
func Int32SliceToUint32Slice(v []int32) []uint32 {
	return *(*[]uint32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// Float32ToUint32 convert float32 to uint32.
func Float32ToUint32(v float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}

// Float32SliceToUint32Slice convert []float32 to []uint32.
func Float32SliceToUint32Slice(v []float32) []uint32 {
	return *(*[]uint32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
