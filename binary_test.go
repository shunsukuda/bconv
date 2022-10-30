package bconv

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBinToU16s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff}}, want: []uint16{0xffff}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff}}, want: []uint16{0xffff, 0xffff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToU16s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToU16s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToU32s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff}}, want: []uint32{0xffffffff}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []uint32{0xffffffff, 0xffffffff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToU32s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToU32s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToU64s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []uint64{0xffffffffffffffff}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []uint64{0xffffffffffffffff, 0xffffffffffffffff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToU64s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToU64s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToI16s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff}}, want: []int16{-1}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff}}, want: []int16{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToI16s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToI16s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToI32s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff}}, want: []int32{-1}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []int32{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToI32s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToI32s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToI64s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: nil},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []int64{-1}},
		{name: "", args: args{b: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}, want: []int64{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToI64s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToI64s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToF32s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0}}, want: nil},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0, 0x0}}, want: []float32{0}},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, want: []float32{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToF32s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToF32s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinToF64s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, want: nil},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, want: []float64{0}},
		{name: "", args: args{b: Binary{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, want: []float64{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToF64s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinToF64s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU16sToBin(t *testing.T) {
	type args struct {
		v []uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: nil}, want: nil},
		{name: "", args: args{v: []uint16{0xffff}}, want: Binary{0xff, 0xff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U16sToBin(tt.args.v); !bytes.Equal(got, tt.want) {
				t.Errorf("U16sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU32sToBin(t *testing.T) {
	type args struct {
		v []uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: nil}, want: nil},
		{name: "", args: args{v: []uint32{0xffffffff}}, want: Binary{0xff, 0xff, 0xff, 0xff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32sToBin(tt.args.v); !bytes.Equal(got, tt.want) {
				t.Errorf("U32sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU64sToBin(t *testing.T) {
	type args struct {
		v []uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: nil}, want: nil},
		{name: "", args: args{v: []uint64{0xffffffffffffffff}}, want: Binary{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U64sToBin(tt.args.v); !bytes.Equal(got, tt.want) {
				t.Errorf("U64sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}
