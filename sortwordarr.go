package piscine

// SortWordArr sorts a slice of strings by their ASCII values in ascending order.
func SortWordArr(words []string) {
	length := len(words)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if words[j] > words[j+1] {
				words[j], words[j+1] = words[j+1], words[j]
			}
		}
	}
}
