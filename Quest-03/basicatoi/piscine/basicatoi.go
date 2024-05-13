package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, digit := range s {
		result = result*10 + int(digit) - '0'
	}
	return result
}
