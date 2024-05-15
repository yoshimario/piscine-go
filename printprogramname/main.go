package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the executable name using os.Args[0]
	exe := os.Args[0]

	// Find the last separator in the path
	sep := '/'
	lastSep := -1
	for i, c := range exe {
		if c == sep {
			lastSep = i
		}
	}

	// Extract the base name from the path
	name := exe
	if lastSep != -1 {
		name = exe[lastSep+1:]
	}

	// Print each character using z01.PrintRune
	for _, r := range name {
		z01.PrintRune(r)
	}

	// Print a newline
	z01.PrintRune('\n')
}
