package bconv

import "encoding/binary"

var BE bigEndian

type bigEndian struct{}

// BinToU16 convert []byte to uint16(big-endian).
func (bigEndian) BinToU16(b Binary) uint16 {
	if len(b) >= 2 {
		return binary.BigEndian.Uint16(b)
		//return uint16(b[1]) | uint16(b[0])<<8
	} else {
		return 0
	}
}

// BinToU32 convert []byte to uint32(big-endian).
func (bigEndian) BinToU32(b Binary) uint32 {
	if len(b) >= 4 {
		return binary.BigEndian.Uint32(b)
		//return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	} else {
		return 0
	}
}

// BinToU64 convert []byte to uint64(big-endian).
func (bigEndian) BinToU64(b Binary) uint64 {
	if len(b) >= 8 {
		return binary.BigEndian.Uint64(b)
		/*
			return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
				uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
		*/
	} else {
		return 0
	}
}

// BinToU16s convert []byte to []uint16(big-endian).
func (bigEndian) BinToU16s(b Binary) []uint16 {
	v := BinToU16s(b)
	for i := range v {
		v[i] = BE.BinToU16(b[:2])
		b = b[2:]
	}
	return v
}

// BinToU32s convert []byte to []uint32(big-endian).
func (bigEndian) BinToU32s(b Binary) []uint32 {
	v := BinToU32s(b)
	for i := range v {
		v[i] = BE.BinToU32(b[:4])
		b = b[4:]
	}
	return v
}

// BinToU64s convert []byte to []uint64(big-endian).
func (bigEndian) BinToU64s(b Binary) []uint64 {
	v := BinToU64s(b)
	for i := range v {
		v[i] = BE.BinToU64(b[:8])
		b = b[8:]
	}
	return v
}

// U16ToBin convert uint16 to []byte(big-endian).
func (bigEndian) U16ToBin(v uint16) Binary {
	b := make(Binary, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
	/*
		return Binary{
			byte(v >> 8), byte(v),
		}
	*/
}

// U32ToBin convert uint32 to []byte(big-endian).
func (bigEndian) U32ToBin(v uint32) Binary {
	b := make(Binary, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
	/*
		return Binary{
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v),
		}
	*/
}

// U64ToBin convert uint64 to []byte(big-endian).
func (bigEndian) U64ToBin(v uint64) Binary {
	b := make(Binary, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
	/*
		return Binary{
			byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32),
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v),
		}
	*/
}

// U16ToBin convert []uint16 to []byte(big-endian).
func (bigEndian) U16sToBin(v []uint16) Binary {
	b := U16sToBin(v)
	for i := range v {
		binary.BigEndian.PutUint16(b[i*2:], v[i])
	}
	return b
}

// U32ToBin convert []uint32 to []byte(big-endian).
func (bigEndian) U32sToBin(v []uint32) Binary {
	b := U32sToBin(v)
	for i := range v {
		binary.BigEndian.PutUint32(b[i*4:], v[i])
	}
	return b
}

// U64ToBin convert []uint64 to []byte(big-endian).
func (bigEndian) U64sToBin(v []uint64) Binary {
	b := U64sToBin(v)
	for i := range v {
		binary.BigEndian.PutUint64(b[i*8:], v[i])
	}
	return b
}
