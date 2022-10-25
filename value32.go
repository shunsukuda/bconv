package bconv

import (
	"reflect"
	"unsafe"
)

// U32ToInt32 convert uint32 to int32.
func U32ToInt32(v uint32) int32 {
	return *(*int32)(unsafe.Pointer(&v))
}

// U32sToI32s convert []uint32 to []int32.
func U32sToI32s(v []uint32) []int32 {
	return *(*[]int32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// U32ToF32 convert uint32 to float32.
func U32ToF32(v uint32) float32 {
	return *(*float32)(unsafe.Pointer(&v))
}

// U32sToF32s convert []uint32 to []float32.
func U32sToF32s(v []uint32) []float32 {
	return *(*[]float32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// I32ToU32 convert int32 to uint32.
func I32ToU32(v int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}

// I32sToU32s convert []int32 to []uint32.
func I32sToU32s(v []int32) []uint32 {
	return *(*[]uint32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}

// F32ToU32 convert float32 to uint32.
func F32ToU32(v float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&v))
}

// F32sToU32s convert []float32 to []uint32.
func F32sToU32s(v []float32) []uint32 {
	return *(*[]uint32)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: (*reflect.SliceHeader)(unsafe.Pointer(&v)).Data,
			Len:  len(v),
			Cap:  len(v),
		}))
}
