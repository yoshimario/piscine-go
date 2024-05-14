package piscine

func Capitalize(s string) string {
	capitalizeNext := true
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			if capitalizeNext {
				result += string(char - 'a' + 'A') // Convert to uppercase
				capitalizeNext = false
			} else {
				result += string(char) // Keep lowercase
			}
		} else if char >= 'A' && char <= 'Z' {
			if capitalizeNext {
				result += string(char) // Keep uppercase
				capitalizeNext = false
			} else {
				result += string(char - 'A' + 'a') // Convert to lowercase
			}
		} else if char >= '0' && char <= '9' {
			result += string(char)
			capitalizeNext = false
		} else {
			result += string(char)
			capitalizeNext = true
		}
	}

	return result
}
