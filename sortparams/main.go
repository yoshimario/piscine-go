package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	// Simple bubble sort implementation
	for pass := 0; pass < len(arg); pass++ {
		for current := 0; current < len(arg)-pass-1; current++ {
			if arg[current] > arg[current+1] {
				arg[current], arg[current+1] = arg[current+1], arg[current]
			}
		}
	}

	// Iterate over the sorted arguments
	for _, argument := range arg {
		// Iterate over each rune in the argument
		for _, char := range argument {
			z01.PrintRune(char)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
