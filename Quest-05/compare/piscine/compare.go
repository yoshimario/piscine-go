package piscine

func Compare(a, b string) int {
	// Loop through the strings comparing characters
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1 // a is less than b
		} else if a[i] > b[i] {
			return 1 // a is greater than b
		}
	}

	// If the loop finishes, the common part of the strings is equal
	// Check the length of the strings to determine the result
	if len(a) < len(b) {
		return -1 // a is shorter than b
	} else if len(a) > len(b) {
		return 1 // a is longer than b
	}

	// If the lengths are equal and no difference is found, return 0
	return 0 // a and b are equal
}
