package bconv

import (
	"bytes"
	"testing"
)

func TestBinToStr(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{b: []byte("abc12345")}, want: "abc12345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinToStr(tt.args.b); got != tt.want {
				t.Errorf("BinToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeBinToStr(t *testing.T) {
	type args struct {
		b Binary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{b: Binary("abc12345")}, want: "abc12345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeBinToStr(tt.args.b); got != tt.want {
				t.Errorf("SafeBinToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToBin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{s: "abc12345"}, want: Binary("abc12345")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToBin(tt.args.s); !bytes.Equal(got, tt.want) {
				t.Errorf("StrToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeStrToBin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Binary
	}{
		{name: "", args: args{s: "abc12345"}, want: Binary("abc12345")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeStrToBin(tt.args.s); !bytes.Equal(got, tt.want) {
				t.Errorf("SafeStrToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
