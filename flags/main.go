package main

import (
	"fmt"
	"os"
)

func main() {
	isOrderFlag, _ := containsFlag(os.Args[1:], "--order", 7, false)
	isInsertFlag, insertValue := containsFlag(os.Args[1:], "--insert", 8, true)
	isHelpFlag, _ := containsFlag(os.Args[1:], "--help", 6, false)
	if isHelpFlag {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("    This flag inserts the string passed as argument.")

		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("    This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}

	otherArgs := getArguments(os.Args[1:])

	for _, arg := range otherArgs {
		var runes []rune
		if isOrderFlag {
			runes = sortRuneTable([]rune(arg))
		} else {
			runes = []rune(arg)
		}
		fmt.Print(string(runes))
	}

	if isInsertFlag {
		if isOrderFlag {
			fmt.Println(string(sortRuneTable([]rune(insertValue))))
		} else {
			fmt.Println(insertValue)
		}
	}
}

func sortRuneTable(table []rune) []rune {
	for currentIndex := 1; currentIndex < stringLength(table); currentIndex++ {
		currentValue := table[currentIndex]
		previousIndex := currentIndex - 1

		for previousIndex >= 0 && table[previousIndex] > currentValue {
			table[previousIndex+1] = table[previousIndex]
			previousIndex = previousIndex - 1
		}
		table[previousIndex+1] = currentValue
	}
	return table
}

func stringLength(s []rune) int {
	count := 0
	for range s {
		count++
	}

	return count
}

func getArguments(args []string) []string {
	var result []string

	for _, arg := range args {
		if !containsRune(arg, '-') {
			result = append(result, arg)
		}
	}

	return result
}

func containsRune(s string, check rune) bool {
	for _, r := range s {
		if r == check {
			return true
		}
	}

	return false
}

func containsFlag(args []string, flagToFind string, flagToFindLen int, takeParam bool) (bool, string) {
	for _, arg := range args {
		if containsRune(arg, '-') {
			if arg[:flagToFindLen] == flagToFind || arg == flagToFind[1:3] {
				if takeParam {
					return true, arg[flagToFindLen+1:]
				} else {
					return true, ""
				}
			}
		}
	}

	return false, ""
}
