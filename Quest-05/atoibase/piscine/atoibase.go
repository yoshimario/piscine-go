package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, char := range s {
		value := indexInBase(char, base)
		if value == -1 {
			return 0
		}
		result = result*baseLen + value
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func indexInBase(char rune, base string) int {
	for i, baseChar := range base {
		if char == baseChar {
			return i
		}
	}
	return -1
}