package piscine

// Any checks if at least one element in the slice returns true when passed through the function f.
func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}