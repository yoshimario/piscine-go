package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if isValidBase(base) {
		printBase(convertToBase(nbr, base))
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

// Function to validate the base
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

// Function to convert number to the specified base
func convertToBase(nbr int, base string) string {
	neg := false
	if nbr < 0 {
		neg = true
		nbr *= -1
	}

	result := ""
	lenBase := len(base)

	for nbr > 0 {
		remainder := nbr % lenBase
		result = string(base[remainder]) + result
		nbr /= lenBase
	}

	if neg {
		result = "-" + result
	}

	return result
}

// Function to print the result
func printBase(result string) {
	for _, char := range result {
		z01.PrintRune(char)
	}
}
