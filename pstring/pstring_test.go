package main

import (
	"strings"
	"testing"
)

func TestHexStringToByte(t *testing.T) {
	var b byte = 0xA0
	if h, _ := HexStringToByte("0xa0"); h != b {
		t.Error("HexStringToByte(\"0xa0\") != 0xA0")
	}
}

func TestHexStringToBytes(t *testing.T) {
	var b byte = 0xA0
	if h := HexStringToBytes("0xa0"); h[0] != b {
		t.Error("HexStringToBytes(\"0xa0\") != 0xA0")
	}
}

func TestIncludesString(t *testing.T) {
	var s = strings.Split("Hello World", " ")
	if !IncludesString(s, "Hello") {
		t.Error("IncludesString(s, \"Hello\") != true")
	}
	if IncludesString(s, "hello") {
		t.Error("IncludesString(s, \"hello\") != false")
	}
}

func TestIncludeKey(t *testing.T) {
	m := make(map[string]interface{})
	m["a"] = 1
	if !IncludeKey(m, "a") {
		t.Error("IncludeKey(m, \"a\") != true")
	}
	if IncludeKey(m, "b") {
		t.Error("IncludeKey(m, \"b\") != false")
	}
}

func TestCleanString(t *testing.T) {
	var s = "Hello\n World"
	if CleanString(s) != "Hello World" {
		t.Errorf("CleanString(s) => %s != \"Hello World\"", CleanString(s))
	}
}
