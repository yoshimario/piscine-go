package piscine

// Map applies function f to each element of the slice a and returns a slice of boolean values.
func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}
