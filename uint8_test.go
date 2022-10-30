package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestU8ToI8(t *testing.T) {
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
			if got := U8ToI8(tt.args.v); got != tt.want {
				t.Errorf("U8ToI8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU8sToI8s(t *testing.T) {
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
			if got := U8sToI8s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U8sToI8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI8ToU8(t *testing.T) {
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
			if got := I8ToU8(tt.args.v); got != tt.want {
				t.Errorf("I8ToU8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI8sToU8s(t *testing.T) {
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
			if got := I8sToU8s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I8sToU8s() = %v, want %v", got, tt.want)
			}
		})
	}
}
