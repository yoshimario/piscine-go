package piscine

// CountIf counts the number of elements in the slice tab for which the function f returns true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}
