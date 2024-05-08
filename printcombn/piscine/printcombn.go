package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	digits := make([]rune, n)

	var printCombNRec func(digits []rune, index int, start rune)
	printCombNRec = func(digits []rune, index int, start rune) {
		if index == n {
			for i, digit := range digits {
				if i > 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(digit)
			}
			return
		}
		for i := start; i <= '9'; i++ {
			digits[index] = i
			printCombNRec(digits, index+1, i+1)
		}
	}

	printCombNRec(digits, 0, '0')
	z01.PrintRune('\n')
}
