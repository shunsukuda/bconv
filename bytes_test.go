package bconv

import (
	"testing"
)

func TestBytesToUint16(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "", args: args{b: nil}, want: 0x0000},
		{name: "", args: args{b: []byte{0x12}}, want: 0x0000},
		{name: "", args: args{b: []byte{0x12, 0x34}}, want: 0x3412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint16(tt.args.b); got != tt.want {
				t.Errorf("BytesToUint16() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestBytesToUint32(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "", args: args{b: nil}, want: 0x0000},
		{name: "", args: args{b: []byte{0x12, 0x34, 0x56}}, want: 0x0000000},
		{name: "", args: args{b: []byte{0x12, 0x34, 0x56, 0x78}}, want: 0x78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint32(tt.args.b); got != tt.want {
				t.Errorf("BytesToUint32() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestBytesToUint64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "", args: args{b: nil}, want: 0x00000000},
		{name: "", args: args{b: []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde}}, want: 0x000000000000000},
		{name: "", args: args{b: []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}}, want: 0xf0debc9a78563412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint64(tt.args.b); got != tt.want {
				t.Errorf("BytesToUint64() = %x, want %x", got, tt.want)
			}
		})
	}
}
