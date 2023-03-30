package pbits

import (
	"fmt"
)

// ByteFromInt Int to byte
func ByteFromInt(i int) byte {
	return byte(i)
}

// ByteToHexString Convert byte to HEX String ie. 0XA0
func ByteToHexString(b byte) string {
	return fmt.Sprintf("0x%02x", b)
}

// PrintByte Print bytes
func PrintByte(b byte) {
	fmt.Printf("hex: 0x%02x\t", b)
	fmt.Printf("uint8: %v\t", b)
	fmt.Printf("bits: %08b\n", b)
}

/*
Byte
Create byte from string

Example:

	Byte("1001") -> 0x09
*/
func Byte(t string) byte {
	var b1 byte = '1'
	l := len(t)
	var b byte = 0 << l
	for i, char := range t {
		if b1 == byte(char) {
			SetBit(&b, uint(l-i)) // backward
		}
	}
	return b
}

/*
SetBit
Set bit at position. First position is 1

Example:

	var b byte = 0x00
	setBit(&b, 1)
	b == 0x01
*/
func SetBit(b *byte, pos uint) {
	*b |= 1 << (pos - 1)
}

/*
	Clear bit at position. First position is 1

Example:

	var b1 byte = 0x01
	fmt.Printf("bits: %08b\n", b1)
	clearBit(&b1, 1)
	fmt.Printf("bits: %08b\n", b1)
*/
func ClearBit(b *byte, pos uint) {
	*b &= ^(1 << (pos - 1))
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
func SliceByte8(b byte, startPosition int, length int) byte {
	left := 8 - (startPosition - 1) - length
	right := 8 - length
	return b << left >> right
}

func Byte8bitsFromInts(a, b int) byte {
	return byte(a<<4 | b)
}
