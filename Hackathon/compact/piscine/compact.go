package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	count := 0
	for _, v := range s {
		if v != "" {
			s[count] = v
			count++
		}
	}
	*ptr = s[:count]
	return count
}
