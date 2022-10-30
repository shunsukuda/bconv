package bconv

import (
	"reflect"
	"testing"
)

func Test_littleEndian_BinToU16(t *testing.T) {
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
		{name: "", args: args{b: Binary{0x12, 0x34}}, want: 0x3412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU16(tt.args.b); got != tt.want {
				t.Errorf("littleEndian.BinToU16() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_BinToU32(t *testing.T) {
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
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: 0x78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU32(tt.args.b); got != tt.want {
				t.Errorf("littleEndian.BinToU32() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_BinToU64(t *testing.T) {
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
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: 0xf0debc9a78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU64(tt.args.b); got != tt.want {
				t.Errorf("littleEndian.BinToU64() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_BinToU16s(t *testing.T) {
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
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78}}, want: []uint16{0x3412, 0x7856}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU16s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.BinToU16s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_BinToU32s(t *testing.T) {
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
		{name: "", args: args{b: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: []uint32{0x78563412, 0xf0debc9a}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU32s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.BinToU32s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_BinToU64s(t *testing.T) {
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
			want: []uint64{0xf0debc9a78563412, 0xf0debc9a78563412}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.BinToU64s(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.BinToU64s() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U16ToBin(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x3412}, want: Binary{0x12, 0x34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U16ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U16ToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U32ToBin(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0x78563412}, want: Binary{0x12, 0x34, 0x56, 0x78}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U32ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U32ToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U64ToBin(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: 0xf0debc9a78563412}, want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U64ToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U64ToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U16sToBin(t *testing.T) {
	type args struct {
		v []uint16
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint16{0x3412, 0x7856}}, want: Binary{0x12, 0x34, 0x56, 0x78}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U16sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U16sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U32sToBin(t *testing.T) {
	type args struct {
		v []uint32
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint32{0x78563412, 0xf0debc9a}}, want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U32sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U32sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_U64sToBin(t *testing.T) {
	type args struct {
		v []uint64
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{v: []uint64{0xf0debc9a78563412, 0xf0debc9a78563412}},
			want: Binary{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LE.U64sToBin(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("littleEndian.U64sToBin() = 0x%x, want 0x%x", got, tt.want)
			}
		})
	}
}
