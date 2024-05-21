package piscine

// AdvancedSortWordArr sorts a slice of strings based on the comparison function compareFunc.
func AdvancedSortWordArr(strings []string, compareFunc func(a, b string) int) {
	length := len(strings)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if compareFunc(strings[j], strings[j+1]) > 0 {
				strings[j], strings[j+1] = strings[j+1], strings[j]
			}
		}
	}
}

// Compare is an example comparison function to sort strings in ascending order by ASCII values.
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
