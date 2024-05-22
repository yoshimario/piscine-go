package piscine

// AdvancedSortWordArr sorts a slice of strings based on the comparison function compareFunc.
func AdvancedSortWordArr(words []string, compareFunc func(a, b string) int) {
	length := len(words)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if compareFunc(words[j], words[j+1]) > 0 {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
}
