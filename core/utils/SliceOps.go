package utils

// FindElement is a generic function to find an element in a slice
// It returns the index of the element if found, else it returns -1
func FindElement[T comparable](slice []T, element T) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}

	return -1
}
