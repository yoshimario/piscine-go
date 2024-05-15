package main

import (
	"fmt"
	"os"
)

func main() {
	commandLineArgs := os.Args
	if len(commandLineArgs) <= 2 {
		switch {
		case len(commandLineArgs) == 1:
			fmt.Print("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.\n")
		case commandLineArgs[1][:2] == "--" || commandLineArgs[0][:2] == "-h" || commandLineArgs[1] == "-h" || commandLineArgs[1] == "-help" || len(commandLineArgs) == 1:
			fmt.Print("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.\n")
		default:
			fmt.Println(commandLineArgs[1])
		}
	}
	if len(commandLineArgs) > 2 {
		switch {
		case len(commandLineArgs[1]) == 2 || len(commandLineArgs[1]) == 7:
			fmt.Println(sortString(commandLineArgs[2]))
		case commandLineArgs[2][:2] == "--" || commandLineArgs[2][:2] == "-o":
			fmt.Println(sortString(insertString(commandLineArgs[1], commandLineArgs[3])))
		default:
			fmt.Println(insertString(commandLineArgs[1], commandLineArgs[2]))
		}
	}
}

func insertString(argument string, valueToInsert string) string {
	if len(argument) > 8 {
		if argument[:9] == "--insert=" {
			return valueToInsert + argument[9:]
		}
		if argument[:3] == "-i=" {
			return valueToInsert + argument[3:]
		}
		return valueToInsert
	} else {
		if argument[:3] == "-i=" {
			return valueToInsert + argument[3:]
		}
	}
	return valueToInsert
}

func sortString(chaos string) string {
	var characters []string
	for _, char := range chaos {
		characters = append(characters, string(char))
	}
	for i := 0; i < len(characters)-1; i++ {
		for j := 0; j < len(characters)-i-1; j++ {
			if characters[j] > characters[j+1] {
				characters[j], characters[j+1] = characters[j+1], characters[j]
			}
		}
	}
	sortedString := ""
	for _, char := range characters {
		sortedString += char
	}
	return sortedString
}
