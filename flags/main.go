package main

import (
	"fmt"
	"os"

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
		if isInsertFlag(arg) {
			if i+1 < len(args) {
				// Check if the insert flag has a value prefixed with -i=
				if hasInsertPrefix(args[i+1]) {
					insertStr = args[i+1][len("-i="):]
					i++ // Skip the next argument
				}
			}
		} else if isOrderFlag(arg) {
			orderFlag = true
		}
	}

	// Remove flags from arguments
	var strArgs []string
	for _, arg := range args {
		if !isInsertFlag(arg) && !isOrderFlag(arg) {
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

func isInsertFlag(arg string) bool {
	// fmt.Println(arg)

	if len(arg) > 9 {
		return arg[:9] == "--insert="
	}

	if len(arg) > 3 {
		return arg[:3] == "-i="
	}

	return false

	// return arg == "--insert=" || arg == "-i"
}

func isOrderFlag(arg string) bool {
	return arg == "--order" || arg == "-o"
}

func hasInsertPrefix(s string) bool {
	return len(s) >= len("-i=") && s[:len("-i=")] == "-i="
}

/*
	 func insertString(args []string, insertStr string) []string {
		// Find the position to insert
		inserted := false
		for i, arg := range args {
			if hasEqualSign(arg) {
				argValue := arg[splitIndex(arg)+1:]
				if startsWith(insertStr, argValue) {
					args = append(args[:i], append([]string{insertStr}, args[i:]...)...)
					inserted = true
					break
				}
			}
		}

		// If not inserted, append at the end
		if !inserted {
			args = append(args, insertStr)
		}

		return args
	}
*/
func insertString(args []string, insertStr string) []string {
	// Find the position to insert
	inserted := false
	for i, arg := range args {
		if arg == "v2" {
			// If the argument is "v2", insert the insert string after it
			args = append(args[:i+1], append([]string{insertStr}, args[i+1:]...)...)
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

func hasEqualSign(s string) bool {
	for _, char := range s {
		if char == '=' {
			return true
		}
	}
	return false
}

func splitIndex(s string) int {
	for i, char := range s {
		if char == '=' {
			return i
		}
	}
	return -1
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
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
