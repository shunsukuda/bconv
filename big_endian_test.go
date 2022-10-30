package bconv

import (
	"reflect"
	"testing"
)

func Test_bigEndian_BinToU16(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "", args: args{b: nil}, want: 0},
		{name: "", args: args{b: Binary{0x12}}, want: 0},
		{name: "", args: args{b: Binary{0x12, 0x34}}, want: 0x1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU16(tt.args.b); got != tt.want {
				t.Errorf("bigEndian.BinToU16() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_BinToU32(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{b: nil}, want: 0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56}}, want: 0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: 0x12345678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU32(tt.args.b); got != tt.want {
				t.Errorf("bigEndian.BinToU32() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_BinToU64(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{b: nil}, want: 0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde}}, want: 0},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: 0x123456789abcdef0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU64(tt.args.b); got != tt.want {
				t.Errorf("bigEndian.BinToU64() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_BinToU16s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0x12}}, want: nil},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: []uint16{0x1234, 0x5678}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU16s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.BinToU16s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_BinToU32s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56}}, want: nil},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: []uint32{0x12345678, 0x9abcdef0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU32s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.BinToU32s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_BinToU64s(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{name: "", args: args{b: nil}, want: nil},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde}}, want: nil},
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
			want: []uint64{0x123456789abcdef0, 0x123456789abcdef0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.BinToU64s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.BinToU64s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U16ToBin(t *testing.T) {
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
			if got := BE.U16ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U16ToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U32ToBin(t *testing.T) {
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
			if got := BE.U32ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U32ToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U64ToBin(t *testing.T) {
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
			if got := BE.U64ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U64ToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U16sToBin(t *testing.T) {
	type args struct {
		v []uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint16{0x1234, 0x5678}}, want: Binary{0x12, 0x34, 0x56, 0x78}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.U16sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U16sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U32sToBin(t *testing.T) {
	type args struct {
		v []uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint32{0x12345678, 0x9abcdef0}}, want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.U32sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U32sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_U64sToBin(t *testing.T) {
	type args struct {
		v []uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint64{0x123456789abcdef0, 0x123456789abcdef0}},
			want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BE.U64sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigEndian.U64sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}
