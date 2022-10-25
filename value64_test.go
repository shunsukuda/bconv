package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestU64ToI64(t *testing.T) {
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
			if got := U64ToI64(tt.args.v); got != tt.want {
				t.Errorf("U64ToI64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU64sToI64s(t *testing.T) {
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
			if got := U64sToI64s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U64sToI64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU64ToF64(t *testing.T) {
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
			if got := U64ToF64(tt.args.v); got != tt.want {
				t.Errorf("U64ToF64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU64sToF64s(t *testing.T) {
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
			if got := U64sToF64s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U64sToF64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI64ToU64(t *testing.T) {
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
			if got := I64ToU64(tt.args.v); got != tt.want {
				t.Errorf("I64ToU64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI64sToU64s(t *testing.T) {
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
			if got := I64sToU64s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I64sToU64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF64ToU64(t *testing.T) {
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
			if got := F64ToU64(tt.args.v); got != tt.want {
				t.Errorf("F64ToU64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF64sToU64s(t *testing.T) {
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
			if got := F64sToU64s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("F64sToU64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
