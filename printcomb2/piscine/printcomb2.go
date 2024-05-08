package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for num1 := '0'; num1 <= '9'; num1++ {
		for j := '0'; j <= '9'; j++ {
			for q := '0'; q <= '9'; q++ {
				for k := '0'; k <= '9'; k++ {
					if num1 == '9' && j == '8' && q == '9' && k == '9' {
						z01.PrintRune(num1)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
					} else if q >= num1 && k > j {
						z01.PrintRune(num1)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else if q > num1 && k >= j {
						z01.PrintRune(num1)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else if q > num1 && k < j {
						z01.PrintRune(num1)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(q)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
