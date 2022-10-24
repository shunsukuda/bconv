package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestUint32ToInt32(t *testing.T) {
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
			if got := Uint32ToInt32(tt.args.v); got != tt.want {
				t.Errorf("Uint32ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32SliceToInt32Slice(t *testing.T) {
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
			if got := Uint32SliceToInt32Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32SliceToInt32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32ToFloat32(t *testing.T) {
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
			if got := Uint32ToFloat32(tt.args.v); got != tt.want {
				t.Errorf("Uint32ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32SliceToFloat32Slice(t *testing.T) {
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
			if got := Uint32SliceToFloat32Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32SliceToFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ToUint32(t *testing.T) {
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
			if got := Int32ToUint32(tt.args.v); got != tt.want {
				t.Errorf("Int32ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32SliceToUint32Slice(t *testing.T) {
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
			if got := Int32SliceToUint32Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32SliceToUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ToUint32(t *testing.T) {
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
			if got := Float32ToUint32(tt.args.v); got != tt.want {
				t.Errorf("Float32ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32SliceToUint32Slice(t *testing.T) {
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
			if got := Float32SliceToUint32Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32SliceToUint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
