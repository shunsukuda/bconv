package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestUint16ToInt16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{name: "", args: args{v: math.MaxUint16}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16ToInt16(tt.args.v); got != tt.want {
				t.Errorf("Uint16ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16SliceToInt16Slice(t *testing.T) {
	type args struct {
		v []uint16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{name: "", args: args{v: []uint16{math.MaxUint16, math.MaxUint16 - 1, math.MaxUint16 - 2}}, want: []int16{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16SliceToInt16Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16SliceToInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16ToUint16(t *testing.T) {
	type args struct {
		v int16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "", args: args{v: -1}, want: math.MaxUint16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16ToUint16(tt.args.v); got != tt.want {
				t.Errorf("Int16ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16SliceToUint16Slice(t *testing.T) {
	type args struct {
		v []int16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{name: "", args: args{v: []int16{-1, -2, -3}}, want: []uint16{math.MaxUint16, math.MaxUint16 - 1, math.MaxUint16 - 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16SliceToUint16Slice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16SliceToUint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
