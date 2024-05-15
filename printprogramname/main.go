package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the executable name using os.Args[0]
	exe := os.Args[0]

	// Extract the base name from the path
	name := filepath.Base(exe)

	// Print each character using z01.PrintRune
	for _, r := range name {
		z01.PrintRune(r)
	}

	// Print a newline
	z01.PrintRune('\n')
}
