package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits [10]int

	// Extract each digit from the number and count their occurrences
	for n > 0 {
		digit := n % 10
		digits[digit]++
		n /= 10
	}

	// Print the digits in ascending order
	for i := 0; i < len(digits); i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune(i) + '0')
		}
	}
}
