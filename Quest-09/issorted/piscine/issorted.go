package piscine

// IsSorted checks if the slice of integers a is sorted in ascending or descending order using the comparison function f.
func IsSorted(f func(a, b int) int, a []int) bool {
	Ascending := true
	Descending := true

	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if cmp > 0 {
			Ascending = false
		} else if cmp < 0 {
			Descending = false
		}
	}

	return Ascending || Descending
}

// MyFunc is a sample comparison function.
func MyFunc(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	}
	return 0
}
