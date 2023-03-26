package pmath

func Version() string {
	return "0.0.1"
}

// BoolToInt Convert bool to int
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
