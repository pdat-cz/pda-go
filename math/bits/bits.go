package bits

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	math2 "github.com/pdat-cz/pda-go/pkg/math"
	"math"
	"strings"
)

const (
	errInvalidBinaryString = "invalid binary string: %s, should be like '01010101'"
	errStringTooLong       = "string %v is too long, maximum is 8 characters"
	errPositionOutOfRange  = "position %v is out of range 1..8"
	ByteSize               = 8
)

// Byte8FromInt Int to byte
func Byte8FromInt(i int) byte {
	return byte(i)
}

// Byte8ToHexString Convert byte to HEX String ie. 0XA0
func Byte8ToHexString(b byte) string {
	return fmt.Sprintf("0x%02x", b)
}

// Byte8ToBinaryString Convert byte to binary string ie. 01010101
func Byte8ToBinaryString(b byte) string {
	return fmt.Sprintf("%08b", b)
}

// PrintByte8 Print bytes
func PrintByte8(b byte) {
	fmt.Printf("hex: 0x%02x\t", b)
	fmt.Printf("uint8: %v\t", b)
	fmt.Printf("bits: %08b\n", b)
}

// Byte8FromString Convert string to byte
//
//	Example:
//	Byte8FromString("1001") -> 0x09
func Byte8FromString(s string) (byte, error) {
	// Check if the string includes only 0 and 1
	if !IsBinaryString(s) {
		return 0, fmt.Errorf(errInvalidBinaryString, s)
	}
	// Check maximum length
	if len(s) > ByteSize {
		return 0, fmt.Errorf(errStringTooLong, s)
	}
	//
	var result byte
	bitPosition := len(s) - 1 // Start from the most significant bit based on the string length
	//
	for _, char := range s {
		if char == '1' {
			result |= 1 << bitPosition
		}
		bitPosition--
	}
	//
	return result, nil
}

// IsBinaryString checks if a string contains only '0' and '1' characters.
func IsBinaryString(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

/*
SetBit8
Set bit at position. First position is 1

Example:

	var b byte = 0x00
	newB, err := SetBit8(b, 1)
	newB = 0x01
*/
func SetBit8(b byte, pos uint) (byte, error) {
	// Check if position is in range
	if pos < 1 || pos > ByteSize {
		return 0, fmt.Errorf(errPositionOutOfRange, pos)
	}
	b |= 1 << (pos - 1)
	return b, nil
}

/*
	Clear bit at position. First position is 1

Example:

	var b1 byte = 0x01
	b2 := clearBit8(b1, 1)
	b2 = 0x00
*/
func ClearBit8(b byte, pos uint) (byte, error) {
	// Check if position is in range
	if pos < 1 || pos > ByteSize {
		return 0, fmt.Errorf(errPositionOutOfRange, pos)
	}
	b &= ^(1 << (pos - 1))
	return b, nil
}

/*
	Is a bit at position ? First position is 1

Example:

	var b byte = 0x01
	fmt.Printf("bits: %08b\n", b)
	fmt.Printf("has a bit at position %v ? %v\n", 1, hasBit(b, 1))
	fmt.Printf("has a bit at position %v ? %v\n", 2, hasBit(b, 2))
*/
func HasBit(b byte, pos uint) bool {
	val := b & (1 << (pos - 1))
	return val > 0
}

// SliceByte8
// Slice decodeFrom 8bit byte
// start is position, first position is 1, length is number of bits
//
//	Example:
//	SliceByte8(byte(0b001101), 2, 2) -> 0b10
func SliceByte8(b byte, startPosition int, length int) (byte, error) {
	// Check if position is in range
	if startPosition < 1 || startPosition > ByteSize {
		return 0, fmt.Errorf(errPositionOutOfRange, startPosition)
	}
	left := ByteSize - (startPosition - 1) - length
	right := ByteSize - length
	return b << left >> right, nil
}

// Byte8bitsFromInts Convert two ints to byte
//
//	Example:
//	Byte8bitsFromInts(0b0011, 0b0101) -> 0b00110101
func Byte8bitsFromInts(a, b int) byte {
	return byte(a<<4 | b)
}

// BytesToHexString ByteTo8Bits Convert byte to 8 bits
func BytesToHexString(b []byte) string {
	var hexStrings []string

	for _, b0 := range b {
		hexStrings = append(hexStrings, Byte8ToHexString(b0))
	}

	return strings.Join(hexStrings, " ")
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

func IntToBytes8(i uint16) []byte {
	b := make([]byte, 1)
	b[0] = byte(i)
	return b
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

func Float32ToBytes32Real(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
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
		if !math2.Contains(y, v) {
			return false
		}
	}

	for _, v := range y {
		if !math2.Contains(x, v) {
			return false
		}
	}

	return true
}

// HexStringToByte Convert HEX string to byte
func HexStringToByte(s string) (byte, error) {
	if s == "" {
		return 0, errors.New("string is empty")
	}
	clean := strings.Replace(s, "0x", "", -1)
	h, _ := hex.DecodeString(clean)
	return h[0], nil
}

// HexStringToBytes Convert HEX string to bytes
func HexStringToBytes(s string) []byte {
	if s == "" {
		return []byte{}
	}
	var bytes []byte
	for _, s0 := range strings.Split(s, " ") {

		h, err := HexStringToByte(s0)
		if err == nil {
			bytes = append(bytes, h)
		}
	}

	return bytes
}
