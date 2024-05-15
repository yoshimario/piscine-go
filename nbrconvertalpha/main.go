package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Check if the first argument is --upper
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	// Convert each argument to its corresponding letter
	for _, arg := range args {
		num := 0
		for _, char := range arg {
			if char < '0' || char > '9' {
				num = -1
				break
			}
			num = num*10 + int(char-'0')
		}

		if num < 1 || num > 26 {
			z01.PrintRune(' ')
			continue
		}

		letter := 'a' + rune(num-1)
		if upper {
			letter -= 32 // Convert to upper case by subtracting 32 from ASCII value
		}
		z01.PrintRune(letter)
	}

	// Print a newline after printing the result
	z01.PrintRune('\n')
}
