package pmath

import "testing"

func TestBoolToInt(t *testing.T) {
	if BoolToInt(true) != 1 {
		t.Error("BoolToInt(true) != 1")
	}
	if BoolToInt(false) != 0 {
		t.Error("BoolToInt(false) != 0")
	}
}
