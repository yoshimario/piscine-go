package piscine

func ToLower(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32) // Convert uppercase to lowercase by adding 32 to ASCII value
		} else {
			result += string(char) // Keep non-uppercase characters unchanged
		}
	}
	return result
}
