package bits

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestByte8FromInt(t *testing.T) {
	check := []struct {
		input int
		want  byte
	}{
		{0, 0x00},
		{1, 0x01},
		{2, 0x02},
		{3, 0x03},
		{250, 0xFA},
		{300, 0x2C},
		{1000000, 0x40},
	}

	for _, c := range check {
		received := Byte8FromInt(c.input)
		if received != c.want {
			t.Errorf("Byte8FromInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}

}

func TestByte8ToHexString(t *testing.T) {
	check := []struct {
		input byte
		want  string
	}{
		{0x00, "0x00"},
		{0x01, "0x01"},
		{0xFF, "0xff"},
	}

	for _, c := range check {
		received := Byte8ToHexString(c.input)
		if received != c.want {
			t.Errorf("Byte8ToHexString(%d) != %s, received: %s", c.input, c.want, received)
		}
	}
}

func TestByte8ToBinaryString(t *testing.T) {
	check := []struct {
		input byte
		want  string
	}{
		{0x00, "00000000"},
		{0x01, "00000001"},
		{0xFF, "11111111"},
	}

	for _, c := range check {
		received := Byte8ToBinaryString(c.input)
		if received != c.want {
			t.Errorf("Byte8ToBinaryString(%d) != %s, received: %s", c.input, c.want, received)
		}
	}
}

func TestByte8FromString(t *testing.T) {
	check := []struct {
		input string
		want  byte
		err   error
	}{
		{"00", 0x00, nil},
		{"01", 0x01, nil},
		{"11111111", 0xFF, nil},
		{"a", 0x00, errors.New(fmt.Sprintf(errInvalidBinaryString, "a"))},
	}

	for _, c := range check {
		received, err := Byte8FromString(c.input)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("Byte8FromString('%s') -> error received: %s != error expected: %s", c.input, err, c.err)
			}
		} else {
			if received != c.want {
				t.Errorf("Byte8FromString(%s) != %d, received: %d", c.input, c.want, received)
			}
		}
	}
}

func TestIsBinaryString(t *testing.T) {
	check := []struct {
		input string
		want  bool
	}{
		{"0x00", false},
		{"0x01", false},
		{"0x", false},
		{"001", true},
		{"11011111", true},
		{"101111111", true},
	}

	for _, c := range check {
		received := IsBinaryString(c.input)
		if received != c.want {
			t.Errorf("IsBinaryString(%s) != %t, received: %t", c.input, c.want, received)
		}
	}
}

func TestSetBit8(t *testing.T) {
	check := []struct {
		input byte
		pos   uint
		want  byte
		err   error
	}{
		{0x00, 1, 0x01, nil},
		{0x00, 2, 0x02, nil},
		{0x00, 3, 0x04, nil},
		{0x00, 4, 0x08, nil},
		{0x00, 9, 0x00, errors.New(fmt.Sprintf(errPositionOutOfRange, 9))},
	}

	for _, c := range check {
		received, err := SetBit8(c.input, c.pos)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("SetBit8(%d,%d) != %d, received: %d", c.input, c.pos, c.err, err)
			}
		} else {
			if received != c.want {
				t.Errorf("SetBit8(%d,%d) != %d, received: %d", c.input, c.pos, c.want, received)
			}
		}
	}
}

func TestClearBit8(t *testing.T) {

	check := []struct {
		input byte
		pos   uint
		want  byte
		err   error
	}{
		{0x01, 1, 0x00, nil},
		{0x02, 2, 0x00, nil},
		{0x04, 3, 0x00, nil},
		{0x08, 4, 0x00, nil},
		{0x00, 9, 0x00, errors.New(fmt.Sprintf(errPositionOutOfRange, 9))},
	}

	for _, c := range check {
		received, err := ClearBit8(c.input, c.pos)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("ClearBit8(%d,%d) != %d, received: %d", c.input, c.pos, c.err, err)
			}
		} else {
			if received != c.want {
				t.Errorf("ClearBit8(%d,%d) != %d, received: %d", c.input, c.pos, c.want, received)
			}
		}
	}
}

func TestHasBit(t *testing.T) {
	check := []struct {
		input byte
		pos   uint
		want  bool
	}{
		{0x01, 1, true},
		{0x01, 2, false},
		{0x02, 2, true},
		{0x04, 3, true},
		{0x08, 4, true},
		{0x00, 9, false},
	}

	for _, c := range check {
		received := HasBit(c.input, c.pos)
		if received != c.want {
			t.Errorf("HasBit(%d,%d) != %t, received: %t", c.input, c.pos, c.want, received)
		}
	}
}

func TestSliceByte8(t *testing.T) {
	check := []struct {
		input  byte
		pos1   int
		length int
		want   byte
		err    error
	}{
		{0x01, 1, 1, 0x01, nil},
		{0x02, 2, 1, 0x01, nil},
	}

	for _, c := range check {
		received, err := SliceByte8(c.input, c.pos1, c.length)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("SliceByte8(%d,%d,%d) != %d, received: %d", c.input, c.pos1, c.length, c.err, err)
			}
		} else {
			if received != c.want {
				t.Errorf("SliceByte8(%d,%d,%d) != %d, received: %d", c.input, c.pos1, c.length, c.want, received)
			}
		}
	}
}

func TestByte8bitsFromInts(t *testing.T) {
	check := []struct {
		input1 string
		input2 string
		want   string
	}{
		{"0011", "1100", "00111100"},
		{"1000", "0001", "10000001"},
	}

	for _, c := range check {
		b1, err := Byte8FromString(c.input1)
		if err != nil {
			t.Errorf("Cannot convert %s to byte", c.input1)
		}
		i1 := int(b1)

		b2, err := Byte8FromString(c.input2)
		if err != nil {
			t.Errorf("Cannot convert %s to byte", c.input2)
		}
		i2 := int(b2)
		//
		received := Byte8bitsFromInts(i1, i2)
		receivedString := Byte8ToBinaryString(received)
		if receivedString != c.want {
			t.Errorf("Byte8bitsFromInts(%s,%s) != %s, received: %s", c.input1, c.input2, c.want, receivedString)
		}
	}
}

func TestBytesToHexString(t *testing.T) {
	check := []struct {
		input []byte
		want  string
	}{
		{[]byte{0x00, 0x01, 0x02}, "0x00 0x01 0x02"},
		{[]byte{0x00, 0x01, 0x02, 0x03}, "0x00 0x01 0x02 0x03"},
	}

	for _, c := range check {
		received := BytesToHexString(c.input)
		if received != c.want {
			t.Errorf("BytesToHexString(%d) != %s, received: %s", c.input, c.want, received)
		}
	}
}

func TestReversedBytes(t *testing.T) {
	check := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{0x00, 0x01, 0x02}, []byte{0x02, 0x01, 0x00}},
		{[]byte{0x00, 0x01, 0x02, 0x03}, []byte{0x03, 0x02, 0x01, 0x00}},
	}

	for _, c := range check {
		received := ReversedBytes(c.input)
		if len(received) != len(c.want) {
			t.Errorf("ReversedBytes(%d) != %d, received: %d", c.input, c.want, received)
		}
		for i, v := range received {
			if v != c.want[i] {
				t.Errorf("ReversedBytes(%d) != %d, received: %d", c.input, c.want, received)
			}
		}
	}
}

func TestBytes8ToInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint16
	}{
		// Min
		{[]byte{0x00}, 0},
		// 1
		{[]byte{0x01}, 1},
		// Max
		{[]byte{0xFF}, 255},
	}

	for _, c := range check {
		received := Bytes8ToInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestIntToBytes8(t *testing.T) {
	check := []struct {
		input uint16
		want  []byte
	}{
		// Min
		{0, []byte{0x00}},
		// 1
		{1, []byte{0x01}},
		// Max
		{255, []byte{0xFF}},
	}

	for _, c := range check {
		received := IntToBytes8(c.input)
		if len(received) != len(c.want) {
			t.Errorf("IntToBytes8(%d) != %d, received: %d", c.input, c.want, received)
		}
		for i, v := range received {
			if v != c.want[i] {
				t.Errorf("IntToBytes8(%d) != %d, received: %d", c.input, c.want, received)
			}
		}
	}
}

func TestBytes16toInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint16
	}{
		// Min
		{[]byte{0x00, 0x00}, 0},
		// 1
		{[]byte{0x01, 0x00}, 1},
		// Max
		{[]byte{0xFF, 0xFF}, 65535},
	}

	for _, c := range check {
		received := Bytes16toInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestBytes24toInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint32
	}{
		// Min
		{[]byte{0x00, 0x00, 0x00}, 0},
		// 1
		{[]byte{0x01, 0x00, 0x00}, 1},
		// Max
		{[]byte{0xFF, 0xFF, 0xFF}, 16777215},
	}

	for _, c := range check {
		received := Bytes24toInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestBytes32toInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint32
	}{
		// Min
		{[]byte{0x00, 0x00, 0x00, 0x00}, 0},
		// Max
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, 4294967295},
	}

	for _, c := range check {
		received := Bytes32toInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestBytes48toInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint64
	}{
		// Min
		{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
		// 1
		{[]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, 1},
		// Max
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 281474976710655},
	}

	for _, c := range check {
		received := Bytes48toInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestBytes64toInt(t *testing.T) {
	check := []struct {
		input []byte
		want  uint64
	}{
		// Min
		{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
		// 1
		{[]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 1},
		// Max
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 18446744073709551615},
	}

	for _, c := range check {
		received := Bytes64toInt(c.input)
		if received != c.want {
			t.Errorf("BytesToInt(%d) != %d, received: %d", c.input, c.want, received)
		}
	}
}

func TestBytes32RealToFloat(t *testing.T) {
	check := []struct {
		input []byte
		want  float32
	}{
		// Min
		{[]byte{0x00, 0x00, 0x00, 0x00}, 0},
		// 1
		{[]byte{0x01, 0x00, 0x00, 0x00}, 1e-45},
		// 256
		{[]byte{0x00, 0x01, 0x00, 0x00}, 3.59e-43},
		// Max
		{[]byte{0xFF, 0xff, 0x7F, 0x7F}, math.MaxFloat32},
	}

	for _, c := range check {
		received := Bytes32RealToFloat(c.input)
		if received != c.want {
			t.Errorf("Bytes32RealToFloat(%d) != %f, received: %f", c.input, c.want, received)
		}
	}
}

func TestFloat32ToBytes32Real(t *testing.T) {
	check := []struct {
		input float32
		want  []byte
	}{
		// Min
		{0, []byte{0x00, 0x00, 0x00, 0x00}},
		// 1
		{1e-45, []byte{0x01, 0x00, 0x00, 0x00}},
		// 256
		{3.59e-43, []byte{0x00, 0x01, 0x00, 0x00}},
		// Max
		{math.MaxFloat32, []byte{0xFF, 0xFF, 0x7F, 0x7F}},
	}

	for _, c := range check {
		received := Float32ToBytes32Real(c.input)
		if len(received) != len(c.want) {
			t.Errorf("Float32ToBytes32Real(%f) != %d, received: %d", c.input, c.want, received)
		}
		for i, v := range received {
			if v != c.want[i] {
				t.Errorf("Float32ToBytes32Real(%f) != %d, received: %d", c.input, c.want, received)
			}
		}
	}
}

func TestBytesAreEqual(t *testing.T) {
	check := []struct {
		input1 []byte
		input2 []byte
	}{
		{[]byte{0x00, 0x01, 0x02}, []byte{0x00, 0x01, 0x02}},
		{[]byte{0x00, 0x01, 0x02, 0x03}, []byte{0x00, 0x01, 0x02, 0x03}},
	}

	for _, c := range check {
		received := BytesAreEqual(c.input1, c.input2)
		if !received {
			t.Errorf("BytesAreEqual(%d,%d) != true, received: %t", c.input1, c.input2, received)
		}
	}
}

func TestHexStringToByte(t *testing.T) {
	check := []struct {
		input string
		want  byte
		err   error
	}{
		{"0x00", byte(0x00), nil},
		{"0x01", byte(0x01), nil},
	}

	for _, c := range check {
		received, err := HexStringToByte(c.input)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf("HexStringToByte('%s') -> error received: %s != error expected: %s", c.input, err, c.err)
			}
		} else {
			if received != c.want {
				t.Errorf("HexStringToByte('%s') -> %d != %d", c.input, c.want, received)
			}
		}
	}
}

func TestHexStringToBytes(t *testing.T) {
	check := []struct {
		input string
		want  []byte
	}{
		{"0x00", []byte{0x00}},
		{"0x01", []byte{0x01}},
		{"0x01 0xFF", []byte{0x01, 0xFF}},
	}

	for _, c := range check {
		received := HexStringToBytes(c.input)
		if len(received) != len(c.want) {
			t.Errorf("HexStringToBytes('%s') -> %d != %d", c.input, c.want, received)
		}
		for i, v := range received {
			if v != c.want[i] {
				t.Errorf("HexStringToBytes('%s') -> %d != %d", c.input, c.want, received)
			}
		}
	}
}
