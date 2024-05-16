package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	wordCount := len(args)
	if wordCount > 0 {
		vowels := "AEIOUaeiou"
		var str, vowelsStr, reversedVowels string

		for i, word := range args {
			str += word
			if i != wordCount-1 {
				str += " "
			}

			for _, char := range word {
				for _, v := range vowels {
					if char == v {
						vowelsStr += string(char)
					}
				}
			}
		}

		vowelsLen := len(vowelsStr)
		for i := vowelsLen - 1; i >= 0; i-- {
			reversedVowels += string(vowelsStr[i])
		}

		var result string
		vowelIndex := 0
		for _, char := range str {
			isVowel := false
			for _, v := range vowels {
				if char == v {
					result += string(reversedVowels[vowelIndex])
					vowelIndex++
					isVowel = true
					break
				}
			}
			if !isVowel {
				result += string(char)
			}
		}

		for _, char := range result {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
