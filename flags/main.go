package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Check if no arguments are provided or if --help flag is given
	if len(args) == 0 || (len(args) == 1 && (args[0] == "--help" || args[0] == "-h")) {
		printUsage()
		return
	}

	// Initialize variables to store insert string and order flag
	var insertStr string
	orderFlag := false

	// Parse command-line flags
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++ // Skip the next argument
			}
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		}
	}

	// Remove flags from arguments
	var strArgs []string
	for _, arg := range args {
		if arg != "--insert" && arg != "-i" && arg != "--order" && arg != "-o" {
			strArgs = append(strArgs, arg)
		}
	}

	// Insert string if --insert flag is provided
	if insertStr != "" {
		strArgs = insertString(strArgs, insertStr)
	}

	// Order string if --order flag is provided
	if orderFlag {
		strArgs = orderString(strArgs)
	}

	// Print the resulting string
	for _, arg := range strArgs {
		for _, char := range arg {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func printUsage() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
func insertString(args []string, insertStr string) []string {
	// Remove the prefix -i= if present
	if strings.HasPrefix(insertStr, "-i=") {
		insertStr = insertStr[len("-i="):]
	}

	// Find the position to insert
	inserted := false
	for i, arg := range args {
		// Split the arg on "=" if it contains it
		argParts := strings.SplitN(arg, "=", 2)
		if len(argParts) > 1 && strings.HasPrefix(insertStr, argParts[1]) {
			args = append(args[:i], append([]string{insertStr}, args[i:]...)...)
			inserted = true
			break
		}
	}

	// If not inserted, append at the end
	if !inserted {
		args = append(args, insertStr)
	}

	return args
}



// splitOnEqual splits the string s on the first occurrence of "="
func splitOnEqual(s string) []string {
	for i := range s {
		if s[i] == '=' {
			return []string{s[:i], s[i+1:]}
		}
	}
	return []string{s}
}


func orderString(args []string) []string {
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	return args
}
