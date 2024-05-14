package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	foundDigit := false

	for _, let := range s {
		if let >= '0' && let <= '9' {
			num = num*10 + int(let-'0')
			foundDigit = true
		} else if let == '-' && !foundDigit {
			sign = -1
		}
	}

	return num * sign
}
