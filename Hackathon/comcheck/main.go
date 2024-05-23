package main

import "os"

func main() {
	alert := false
	// Iterate through command-line arguments starting from index 1
	for _, arg := range os.Args[1:] {
		// Check if the argument matches any of the specified strings
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			// If a match is found, set the alert flag to true
			alert = true
			break // Exit the loop early since the condition is met
		}
	}
	// If the alert flag is true, print "Alert!!!"
	if alert {
		os.Stdout.WriteString("Alert!!!\n")
	}
}
