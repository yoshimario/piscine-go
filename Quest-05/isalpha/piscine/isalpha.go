package piscine

func IsAlpha(s string) bool {
	// Iterate over each character in the input string
	for _, char := range s {
		// Check if the character is not alphanumeric
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	// If all characters are alphanumeric or the string is empty, return true
	return true
}
