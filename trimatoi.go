package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1

	for _, let := range s {
		if let >= '0' && let <= '9' {
			num = num*10 + int(let-'0')
		} else if let == '-' {
			sign = -1
		}
	}

	return num * sign
}
