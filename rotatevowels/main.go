package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	for _, arg := range os.Args[1:] {
		wordRunes := []rune(arg)
		leftIndex, rightIndex := 0, len(wordRunes)-1
		for leftIndex < rightIndex {
			if !isVowel(wordRunes[leftIndex]) {
				leftIndex++
				continue
			}
			if !isVowel(wordRunes[rightIndex]) {
				rightIndex--
				continue
			}
			wordRunes[leftIndex], wordRunes[rightIndex] = wordRunes[rightIndex], wordRunes[leftIndex]
			leftIndex++
			rightIndex--
		}
		for _, r := range wordRunes {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
