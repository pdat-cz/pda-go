package math

// Contains returns true if the element is in the elements
func Contains[T comparable](elements []T, element T) bool {
	for _, x := range elements {
		if x == element {
			return true
		}
	}
	return false
}
