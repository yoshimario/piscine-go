package piscine

func FirstRune(s string) rune {
	if len(s) > 0 {
		return []rune(s)[0]
	}
	return 0 // Return 0 if the string is empty
}