package bconv

import (
	"math"
	"reflect"
	"testing"
)

func TestU16ToI16(t *testing.T) {
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
			if got := U16ToI16(tt.args.v); got != tt.want {
				t.Errorf("U16ToI16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU16sToI16s(t *testing.T) {
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
			if got := U16sToI16s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U16sToI16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI16ToU16(t *testing.T) {
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
			if got := I16ToU16(tt.args.v); got != tt.want {
				t.Errorf("I16ToU16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI16sToU16s(t *testing.T) {
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
			if got := I16sToU16s(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I16sToU16s() = %v, want %v", got, tt.want)
			}
		})
	}
}
