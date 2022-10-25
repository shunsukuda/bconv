package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestU32ToInt32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "", args: args{v: math.MaxUint32}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32ToInt32(tt.args.v); got != tt.want {
				t.Errorf("U32ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU32sToI32s(t *testing.T) {
	type args struct {
		v []uint32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{name: "", args: args{v: []uint32{math.MaxUint32, math.MaxUint32 - 1, math.MaxUint32 - 2}}, want: []int32{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32sToI32s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U32sToI32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU32ToF32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "", args: args{v: 1}, want: 1e-45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32ToF32(tt.args.v); got != tt.want {
				t.Errorf("U32ToF32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU32sToF32s(t *testing.T) {
	type args struct {
		v []uint32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{name: "", args: args{v: []uint32{1, 2, 3}}, want: []float32{1e-45, 3e-45, 4e-45}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32sToF32s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U32sToF32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI32ToU32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{v: -1}, want: math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I32ToU32(tt.args.v); got != tt.want {
				t.Errorf("I32ToU32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI32sToU32s(t *testing.T) {
	type args struct {
		v []int32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{name: "", args: args{v: []int32{-1, -2, -3}}, want: []uint32{math.MaxUint32, math.MaxUint32 - 1, math.MaxUint32 - 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I32sToU32s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I32sToU32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF32ToU32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{v: 1e-45}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := F32ToU32(tt.args.v); got != tt.want {
				t.Errorf("F32ToU32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF32sToU32s(t *testing.T) {
	type args struct {
		v []float32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{name: "", args: args{v: []float32{1e-45, 3e-45, 4e-45}}, want: []uint32{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := F32sToU32s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("F32sToU32s() = %v, want %v", got, tt.want)
			}
		})
	}
}
