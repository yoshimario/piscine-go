package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the arguments from the command line (excluding the program name)
	args := os.Args[1:]

	// Iterate over the arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		arg := args[i]

		// Iterate over each rune in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}

		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
