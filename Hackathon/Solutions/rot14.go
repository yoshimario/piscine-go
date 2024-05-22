package piscine

func Rot14(s string) string {
	result := make([]rune, len(s))

	for i, c := range s {
		switch {
		case 'a' <= c && c <= 'z':
			result[i] = 'a' + (c-'a'+14)%26
		case 'A' <= c && c <= 'Z':
			result[i] = 'A' + (c-'A'+14)%26
		default:
			result[i] = c
		}
	}

	return string(result)
}
