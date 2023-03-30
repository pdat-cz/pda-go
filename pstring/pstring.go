package pstring

import (
	"encoding/hex"
	"errors"
	"strings"
)

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

// IncludesString If string is in []string, then return true
func IncludesString(m []string, s string) bool {
	for _, x := range m {
		if x == s {
			return true
		}
	}
	return false
}

func IncludeKey(m map[string]interface{}, s string) bool {
	for key := range m {
		if key == s {
			return true
		}
	}
	return false
}

// CleanString remove all unwanted code
func CleanString(s string) string {
	var replacer = strings.NewReplacer(
		"\r\n", "",
		"\r", "",
		"\n", "",
		"\v", "",
		"\f", "",
		"\u0008", "",
		"\u0085", "",
		"\u2028", "",
		"\u2029", "",
	)

	return replacer.Replace(s)
}
