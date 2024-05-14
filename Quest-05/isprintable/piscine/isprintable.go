package piscine

import "unicode"

func IsPrintable(s string) bool {
	for _, char := range s {
		if !unicode.IsPrint(char) {
			return false
		}
	}
	return true
}
