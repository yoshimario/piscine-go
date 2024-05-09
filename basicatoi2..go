package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0 // Return 0 if non-digit character is found
		}
		result = result*10 + int(digit) - '0'
	}
	return result
}
