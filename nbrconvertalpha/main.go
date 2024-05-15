package main

import (
	"os"
	"strconv"

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
		num, err := strconv.Atoi(arg)
		if err != nil || num < 1 || num > 26 {
			z01.PrintRune(' ')
			continue
		}

		letter := 'a' + rune(num-1)
		if upper {
			letter -= 32 // Convert to upper case by subtracting 32 from ASCII value
		}
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
