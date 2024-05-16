package piscine

import "github.com/01-edu/z01"

// PrintWordsTables prints each element of the slice on a separate line
func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
