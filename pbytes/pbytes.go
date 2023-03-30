package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

// BytesToHexString ByteTo8Bits Convert byte to 8 bits
func BytesToHexString(b []byte) string {
	var s string

	for _, b0 := range b {
		if s != "" {
			s = fmt.Sprintf("%s 0x%02x", s, b0)
		} else {
			s = fmt.Sprintf("0x%02x", b0)
		}
	}
	return s
}

// ReversedBytes Reverse bytes
// Example: 0x01 0x02 0x03 -> 0x03 0x02 0x01
func ReversedBytes(s []byte) []byte {
	var r []byte
	// backward
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func Bytes8ToInt(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func Bytes16toInt(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func Bytes24toInt(b []byte) uint32 {
	//,Append change source, so  i am going by this per parts way
	bytes32 := []byte{b[0], b[1], b[2], 0x00}
	return binary.LittleEndian.Uint32(bytes32)
}

func Bytes32toInt(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func Bytes48toInt(b []byte) uint64 {
	bytes64 := []byte{b[0], b[1], b[2], b[3], b[4], b[5], 0x00, 0x00}
	return binary.LittleEndian.Uint64(bytes64)
}

func Bytes64toInt(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Bytes32RealToFloat(b []byte) float32 {
	data := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(data)
}

// BytesAreEqual Are []bytes equally ? Position can be different
func BytesAreEqual(x []byte, y []byte) bool {

	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {

		return false
	}
	if len(x) != len(y) {

		return false
	}

	for _, v := range x {
		if !IncludesByte(y, v) {
			return false
		}
	}

	for _, v := range y {
		if !IncludesByte(x, v) {
			return false
		}
	}

	return true
}

// IncludesByte If byte is in []byte, then return true
func IncludesByte(m []byte, b byte) bool {
	for _, x := range m {
		if x == b {
			return true
		}
	}
	return false
}
