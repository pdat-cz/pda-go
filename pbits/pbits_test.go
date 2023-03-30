package main

import (
	"testing"
)

func TestByteToHexString(t *testing.T) {
	var b byte = 0xA0
	if ByteToHexString(b) != "0xa0" {
		t.Error("ByteToHexString(b) != 0xa0")
	}
}

func TestSetBit(t *testing.T) {
	var b byte = 0x00
	SetBit(&b, 1)
	if b != 0x01 {
		t.Error("SetBit(b, 0) != 0x01")
		PrintByte(b)
	}
	SetBit(&b, 5)
	if b != 0x11 {
		t.Error("SetBit(b, 5) != 0x21")
		PrintByte(b)
	}
}

func TestClearBit(t *testing.T) {

	var b byte = 0x01
	ClearBit(&b, 1)
	if b != 0x00 {
		t.Error("ClearBit(b, 1) != 0x00")
		PrintByte(b)
	}

}

func TestHasBit(t *testing.T) {
	var b byte = 0x01
	if !HasBit(b, 1) {
		t.Error("HasBit(b, 1) != true")
	}
	if HasBit(b, 2) {
		t.Error("HasBit(b, 2) != false")
	}
}

func TestSliceByte8(t *testing.T) {
	var b byte = 0b001001

	if SliceByte8(b, 1, 2) != 0x01 {
		t.Error("SliceByte8(b)[0] != 0x01")
		PrintByte(SliceByte8(b, 1, 2))
	}
	if SliceByte8(b, 2, 3) != 0b100 {
		t.Error("SliceByte8(b)[1] != 0b100")
		PrintByte(SliceByte8(b, 2, 3))
	}
}

func TestByte8bitsFromInts(t *testing.T) {
	var a int = 1
	var b int = 2
	out := Byte("00010010")
	calc := Byte8bitsFromInts(a, b)

	if calc != out {
		t.Errorf("Byte8bitsFromInts(%d,%d) != %s", a, b, ByteToHexString(out))
		PrintByte(calc)
		PrintByte(ByteFromInt(a))
		PrintByte(ByteFromInt(b))
	}
}
