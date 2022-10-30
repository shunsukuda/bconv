package bconv

import "encoding/binary"

var LE littleEndian

type littleEndian struct{}

// BinToU16 convert []byte to uint16(little-endian).
func (littleEndian) BinToU16(b Binary) uint16 {
	if len(b) >= 2 {
		return binary.LittleEndian.Uint16(b)
		//return uint16(b[1])<<8 | uint16(b[0])
	} else {
		return 0
	}
}

// BinToU32 convert []byte to uint32(little-endian).
func (littleEndian) BinToU32(b Binary) uint32 {
	if len(b) >= 4 {
		return binary.LittleEndian.Uint32(b)
		// return uint32(b[3])<<24 | uint32(b[2])<<16 | uint32(b[1])<<8 | uint32(b[0])
	} else {
		return 0
	}
}

// BinToU64 convert []byte to uint64(little-endian).
func (littleEndian) BinToU64(b Binary) uint64 {
	if len(b) >= 8 {
		return binary.LittleEndian.Uint64(b)
		/*
			return uint64(b[7])<<56 | uint64(b[6])<<48 | uint64(b[5])<<40 | uint64(b[4])<<32 |
				uint64(b[3])<<24 | uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
		*/
	} else {
		return 0
	}
}

// BinToU16s convert []byte to []uint16(little-endian).
func (littleEndian) BinToU16s(b Binary) []uint16 {
	v := BinToU16s(b)
	for i := range v {
		v[i] = LE.BinToU16(b[:2])
		b = b[2:]
	}
	return v
}

// BinToU32s convert []byte to []uint32(little-endian).
func (littleEndian) BinToU32s(b Binary) []uint32 {
	v := BinToU32s(b)
	for i := range v {
		v[i] = LE.BinToU32(b[:4])
		b = b[4:]
	}
	return v
}

// BinToU64s convert []byte to []uint64(little-endian).
func (littleEndian) BinToU64s(b Binary) []uint64 {
	v := BinToU64s(b)
	for i := range v {
		v[i] = LE.BinToU64(b[:8])
		b = b[8:]
	}
	return v
}

// U16ToBin convert uint16 to []byte(little-endian).
func (littleEndian) U16ToBin(v uint16) Binary {
	b := make(Binary, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
	/*
		return Binary{
			byte(v), byte(v >> 8),
		}
	*/
}

// U32ToBin convert uint32 to []byte(little-endian).
func (littleEndian) U32ToBin(v uint32) Binary {
	b := make(Binary, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
	/*
		return Binary{
			byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24),
		}
	*/
}

// U64ToBin convert uint64 to []byte(little-endian).
func (littleEndian) U64ToBin(v uint64) Binary {
	b := make(Binary, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
	/*
		return Binary{
			byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24),
			byte(v >> 32), byte(v >> 40), byte(v >> 48), byte(v >> 56),
		}
	*/
}

// U16ToBin convert []uint16 to []byte(little-endian).
func (littleEndian) U16sToBin(v []uint16) Binary {
	b := U16sToBin(v)
	for i := range v {
		binary.LittleEndian.PutUint16(b[i*2:], v[i])
	}
	return b
}

// U32ToBin convert []uint32 to []byte(little-endian).
func (littleEndian) U32sToBin(v []uint32) Binary {
	b := U32sToBin(v)
	for i := range v {
		binary.LittleEndian.PutUint32(b[i*4:], v[i])
	}
	return b
}

// U64ToBin convert []uint64 to []byte(little-endian).
func (littleEndian) U64sToBin(v []uint64) Binary {
	b := U64sToBin(v)
	for i := range v {
		binary.LittleEndian.PutUint64(b[i*8:], v[i])
	}
	return b
}
