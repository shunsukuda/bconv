package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestUint8ToInt8(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "", args: args{v: math.MaxUint8}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8ToInt8(tt.args.v); got != tt.want {
				t.Errorf("Uint8ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8SliceToInt8Slice(t *testing.T) {
	type args struct {
		v []uint8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{name: "", args: args{v: []uint8{math.MaxUint8, math.MaxUint8 - 1, math.MaxUint8 - 2}}, want: []int8{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8SliceToInt8Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8SliceToInt8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8ToUint8(t *testing.T) {
	type args struct {
		v int8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{name: "", args: args{v: -1}, want: math.MaxUint8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8ToUint8(tt.args.v); got != tt.want {
				t.Errorf("Int8ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8SliceToUint8Slice(t *testing.T) {
	type args struct {
		v []int8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{name: "", args: args{v: []int8{-1, -2, -3}}, want: []uint8{math.MaxUint8, math.MaxUint8 - 1, math.MaxUint8 - 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8SliceToUint8Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8SliceToUint8Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
