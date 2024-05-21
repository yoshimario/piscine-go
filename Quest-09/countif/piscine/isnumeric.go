package piscine

func IsNumeric(s string) bool {
	// Iterate over each character in the input string
	for _, char := range s {
		// Check if the character is not a numerical character
		if char < '0' || char > '9' {
			return false
		}
	}
	// If all characters are numerical, return true
	return true
}
