package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for num1 := '0'; num1 <= '9'; num1++ {
		for num2 := '0'; num2 <= '9'; num2++ {
			for num3 := '0'; num3 <= '9'; num3++ {
				if num1 < num2 && num2 < num3 {
					z01.PrintRune(num1)
					z01.PrintRune(num2)
					z01.PrintRune(num3)

					if num1 != '7' || num2 != '8' || num3 != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
