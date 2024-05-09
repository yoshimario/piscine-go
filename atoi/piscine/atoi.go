package piscine

func Atoi(s string) int {
	result := 0
	sign := 1 // Initialize sign to positive

	// Check if the first character is a sign (+ or -)
	if s[0] == '-' {
		sign = -1 // Update sign to negative if the first character is '-'
		s = s[1:] // Remove the sign from the string
	} else if s[0] == '+' {
		s = s[1:] // Remove the sign from the string
	}

	// Convert the remaining string to an integer
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0 // Return 0 if non-digit character is found
		}
		result = result*10 + int(digit-'0')
	}

	// Multiply the result by the sign
	result *= sign

	return result
}
