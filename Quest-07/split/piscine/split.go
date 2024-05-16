package piscine

// Split splits a string s by the separator sep and returns a slice of strings
func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	start := 0
	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}
