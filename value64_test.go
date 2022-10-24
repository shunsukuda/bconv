package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestUint64ToInt64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "", args: args{v: math.MaxUint64}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ToInt64(tt.args.v); got != tt.want {
				t.Errorf("Uint64ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64SliceToInt64Slice(t *testing.T) {
	type args struct {
		v []uint64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "", args: args{v: []uint64{math.MaxUint64, math.MaxUint64 - 1, math.MaxUint64 - 2}}, want: []int64{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64SliceToInt64Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64SliceToInt64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ToFloat64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "", args: args{v: 1}, want: 5e-324},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ToFloat64(tt.args.v); got != tt.want {
				t.Errorf("Uint64ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64SliceToFloat64Slice(t *testing.T) {
	type args struct {
		v []uint64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "", args: args{v: []uint64{1, 2, 3}}, want: []float64{5e-324, 1e-323, 1.5e-323}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64SliceToFloat64Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64SliceToFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToUint64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{v: -1}, want: math.MaxUint64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToUint64(tt.args.v); got != tt.want {
				t.Errorf("Int64ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64SliceToUint64Slice(t *testing.T) {
	type args struct {
		v []int64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{name: "", args: args{v: []int64{-1, -2, -3}}, want: []uint64{math.MaxUint64, math.MaxUint64 - 1, math.MaxUint64 - 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64SliceToUint64Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64SliceToUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ToUint64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{v: 5e-324}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ToUint64(tt.args.v); got != tt.want {
				t.Errorf("Float64ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64SliceToUint64Slice(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{name: "", args: args{v: []float64{5e-324, 1e-323, 1.5e-323}}, want: []uint64{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64SliceToUint64Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64SliceToUint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
