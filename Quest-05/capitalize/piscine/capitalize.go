package piscine

func Capitalize(s string) string {
	// Initialize a flag to indicate if the next character should be capitalized
	capitalizeNext := true

	// Initialize an empty string to store the result
	result := ""

	// Loop through each character in the input string
	for _, char := range s {
		// Check if the current character is a letter or a number
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			// Capitalize the current character if needed
			if capitalizeNext {
				result += string(char)
				capitalizeNext = false
			} else {
				result += string(char)
			}
		} else {
			// If the character is not a letter or a number, reset the flag
			// to capitalize the next character
			result += string(char)
			capitalizeNext = true
		}
	}

	// Return the final result
	return result
}
