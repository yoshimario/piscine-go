package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Iterate over the command line arguments starting from the first argument (ignoring the program name)
	for _, arg := range os.Args[1:] {
		// Iterate over each rune in the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
