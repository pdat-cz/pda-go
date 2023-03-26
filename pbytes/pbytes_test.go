package pbytes

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

func TestBytesToHexString(t *testing.T) {
	var b []byte = []byte{0x01, 0x02, 0x03}
	var out string = "0x01 0x02 0x03"
	var calc string = BytesToHexString(b)
	if calc != out {
		t.Errorf("BytesToHexString(b)=> %s != %s", BytesToHexString(b), out)
	}
}

func TestReversedBytes(t *testing.T) {
	var b []byte = []byte{0x01, 0x02, 0x03}
	var out []byte = []byte{0x03, 0x02, 0x01}
	if bytes.Compare(ReversedBytes(b), out) != 0 {
		t.Error("ReversedBytes(b) != out")
	}
}

func TestBytes8ToInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00}
	var out uint16 = 1
	if Bytes8ToInt(b) != out {
		t.Errorf("Bytes8ToInt(b) => %v != %v", Bytes8ToInt(b), out)
	}
	b = []byte{0x00, 0x01}
	out = 256
	if Bytes8ToInt(b) != out {
		t.Errorf("Bytes8ToInt(b) => %v != %v", Bytes8ToInt(b), out)
	}
}

func TestBytes16toInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00}
	var out uint16 = 1
	if Bytes16toInt(b) != out {
		t.Errorf("TestBytes16toInt(b) => %v != %v", Bytes16toInt(b), out)

	}
	b = []byte{0x00, 0x01}
	out = 256.0
	if Bytes16toInt(b) != out {
		t.Errorf("TestBytes16toInt(b) => %v != %v", Bytes16toInt(b), out)

	}
}

func TestBytes24toInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00, 0x00}
	var out uint32 = 1
	if Bytes24toInt(b) != out {
		t.Errorf("TestBytes24toInt(b) => %v != %v", Bytes24toInt(b), out)

	}
	b = []byte{0x00, 0x01, 0x00}
	out = 256.0
	if Bytes24toInt(b) != out {
		t.Errorf("TestBytes24toInt(b) => %v != %v", Bytes24toInt(b), out)

	}
}

func TestBytes32toInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00, 0x00, 0x00}
	var out uint32 = 1
	if Bytes32toInt(b) != out {
		t.Errorf("TestBytes32toInt(b) => %v != %v", Bytes32toInt(b), out)

	}
	b = []byte{0x00, 0x01, 0x00, 0x00}
	out = 256.0
	if Bytes32toInt(b) != out {
		t.Errorf("TestBytes32toInt(b) => %v != %v", Bytes32toInt(b), out)

	}
}

func TestBytes48toInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00}
	var out uint64 = 1
	if Bytes48toInt(b) != out {
		t.Errorf("TestBytes48toInt(b) => %v != %v", Bytes48toInt(b), out)

	}
	b = []byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x00}
	out = 256.0
	if Bytes48toInt(b) != out {
		t.Errorf("TestBytes48toInt(b) => %v != %v", Bytes48toInt(b), out)

	}
}

func TestBytes64toInt(t *testing.T) {
	var b []byte = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	var out uint64 = 1
	if Bytes64toInt(b) != out {
		t.Errorf("TestBytes64toInt(b) => %v != %v", Bytes64toInt(b), out)

	}
	b = []byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	out = 256.0
	if Bytes64toInt(b) != out {
		t.Errorf("TestBytes64toInt(b) => %v != %v", Bytes64toInt(b), out)

	}
}

func TestBytes32RealToFloat(t *testing.T) {
	var b1 uint32 = math.Float32bits(1.0)
	var b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, b1)
	var out float32 = 1.0
	if Bytes32RealToFloat(b) != out {
		t.Errorf("TestBytes32RealToFloat(b) => %v != %v", Bytes32RealToFloat(b), out)

	}
	b1 = math.Float32bits(1.5)
	binary.LittleEndian.PutUint32(b, b1)
	out = 1.5
	if Bytes32RealToFloat(b) != out {
		t.Errorf("TestBytes32RealToFloat(b) => %v != %v", Bytes32RealToFloat(b), out)
	}
}

func TestBytesAreEqual(t *testing.T) {
	var b []byte = []byte{0x01, 0x02, 0x03}
	var out []byte = []byte{0x01, 0x02, 0x03}
	if !BytesAreEqual(b, out) {
		t.Error("BytesAreEqual(b, out) != true")
	}
}

func TestIncludesByte(t *testing.T) {
	var b []byte = []byte{0x01, 0x02, 0x03}
	if !IncludesByte(b, 0x01) {
		t.Error("IncludesByte(b, 0x01) != true")
	}
	if !IncludesByte(b, 0x02) {
		t.Error("IncludesByte(b, 0x02) != true")
	}
	if !IncludesByte(b, 0x03) {
		t.Error("IncludesByte(b, 0x03) != true")
	}
	if IncludesByte(b, 0x04) {
		t.Error("IncludesByte(b, 0x04) != false")
	}
}
