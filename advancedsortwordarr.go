package piscine

// CompareByASCII compares two strings by their ASCII values.
// It returns a positive number if a > b, a negative number if a < b, and 0 if they are equal.
func CompareByASCII(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
