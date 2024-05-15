package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the executable name
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract the base name from the path
	name := filepath.Base(exe)

	// Print the name of the executable
	fmt.Println(name)
}
