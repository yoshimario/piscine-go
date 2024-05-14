package piscine

func ToLower(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char - 32) // Convert lowercase to uppercase by subtracting 32 from ASCII value
		} else {
			result += string(char) // Keep non-lowercase characters unchanged
		}
	}
	return result
}
