package bconv

import (
	"reflect"
	"testing"
)

func TestBinLEToU16(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34}}, want: 0x3412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinLEToU16(tt.args.b); got != tt.want {
				t.Errorf("BinLEToU16() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinBEToU16(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34}}, want: 0x1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinBEToU16(tt.args.b); got != tt.want {
				t.Errorf("BinBEToU16() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinLEToU32(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: 0x78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinLEToU32(tt.args.b); got != tt.want {
				t.Errorf("BinLEToU32() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinBEToU32(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: 0x12345678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinBEToU32(tt.args.b); got != tt.want {
				t.Errorf("BinBEToU32() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinLEToU64(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: 0xf0debc9a78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinLEToU64(tt.args.b); got != tt.want {
				t.Errorf("BinLEToU64() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestBinBEToU64(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{b: nil}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde}}, want: 0x0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: 0x123456789abcdef0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinBEToU64(tt.args.b); got != tt.want {
				t.Errorf("BinBEToU64() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

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
				t.Errorf("BinToI16s() = %v, want %v", got, tt.want)
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
				t.Errorf("BinToI32s() = %v, want %v", got, tt.want)
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
				t.Errorf("BinToI64s() = %v, want %v", got, tt.want)
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
				t.Errorf("BinToF32s() = %v, want %v", got, tt.want)
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
				t.Errorf("BinToF64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestU16ToBinLE(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x1234}, want: Binary{0x34, 0x12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U16ToBinLE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U16ToBinLE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU16ToBinBE(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x1234}, want: Binary{0x12, 0x34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U16ToBinBE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U16ToBinBE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU32ToBinLE(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x12345678}, want: Binary{0x78, 0x56, 0x34, 0x12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32ToBinLE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U32ToBinLE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU32ToBinBE(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x12345678}, want: Binary{0x12, 0x34, 0x56, 0x78}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U32ToBinBE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U32ToBinBE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU64ToBinLE(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x123456789abcdef0}, want: Binary{0xf0, 0xde, 0xbc, 0x9a, 0x78, 0x56, 0x34, 0x12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U64ToBinLE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U64ToBinLE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func TestU64ToBinBE(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x123456789abcdef0}, want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := U64ToBinBE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("U64ToBinBE() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}
