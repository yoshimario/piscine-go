package piscine

func IsAlpha(s string) bool {
    // Iterate over each character in the input string
    for _, char := range s {
        // Check if the character is not alphanumeric
        if !isAlphaNumeric(char) {
            return false
        }
    }
    // If all characters are alphanumeric or the string is empty, return true
    return true
}

// Helper function to check if a character is alphanumeric
func isAlphaNumeric(char rune) bool {
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
