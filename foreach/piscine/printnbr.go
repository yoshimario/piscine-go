package piscine

import "github.com/01-edu/z01"

// PrintNbr prints an integer using z01.PrintRune.
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	for n > 0 {
		digits = append([]rune{rune(n%10 + '0')}, digits...)
		n /= 10
	}
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
