package bconv_test

import (
	"encoding/binary"
	"testing"

	"github.com/shunsukuda/bconv"
)

func BenchmarkBinLEToU16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinLEToU16([]byte{byte(i), byte(i)})
	}
}

func BenchmarkGoBinLEToU16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.LittleEndian.Uint16([]byte{byte(i), byte(i)})
	}
}
func BenchmarkBinBEToU16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinBEToU16([]byte{byte(i), byte(i)})
	}
}

func BenchmarkGoBinBEToU16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.BigEndian.Uint16([]byte{byte(i), byte(i)})
	}
}

func BenchmarkBinLEToU32(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinLEToU32([]byte{byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkGoBinLEToU32(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.LittleEndian.Uint32([]byte{byte(i), byte(i), byte(i), byte(i)})
	}
}
func BenchmarkBinBEToU32(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinBEToU32([]byte{byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkGoBinBEToU32(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.BigEndian.Uint32([]byte{byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkBinLEToU64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinLEToU64([]byte{byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkGoBinLEToU64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.LittleEndian.Uint64([]byte{byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i)})
	}
}
func BenchmarkBinBEToU64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.BinBEToU64([]byte{byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkGoBinBEToU64(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binary.BigEndian.Uint64([]byte{byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i), byte(i)})
	}
}

func BenchmarkU16ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U16ToBinLE(uint16(i))
	}
}

func BenchmarkGoU16ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 2)
		binary.LittleEndian.PutUint16(s, uint16(i))
	}
}

func BenchmarkU16ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U16ToBinBE(uint16(i))
	}
}

func BenchmarkGoU16ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 2)
		binary.BigEndian.PutUint16(s, uint16(i))
	}
}

func BenchmarkU32ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U32ToBinLE(uint32(i))
	}
}

func BenchmarkGoU32ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 4)
		binary.LittleEndian.PutUint32(s, uint32(i))
	}
}

func BenchmarkU32ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U32ToBinBE(uint32(i))
	}
}

func BenchmarkGoU32ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 4)
		binary.BigEndian.PutUint32(s, uint32(i))
	}
}

func BenchmarkU64ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U64ToBinLE(uint64(i))
	}
}

func BenchmarkGoU64ToBinLE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 8)
		binary.LittleEndian.PutUint64(s, uint64(i))
	}
}

func BenchmarkU64ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bconv.U64ToBinBE(uint64(i))
	}
}

func BenchmarkGoU64ToBinBE(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 8)
		binary.BigEndian.PutUint64(s, uint64(i))
	}
}
