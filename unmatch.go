package piscine

func Unmatch(a []int) int {
	// Create a map to store the frequency of each element
	freqMap := make(map[int]int)

	// Count the frequency of each element in the slice
	for _, num := range a {
		freqMap[num]++
	}

	// Find the element that does not have a corresponding pair
	for num, freq := range freqMap {
		if freq%2 != 0 {
			return num
		}
	}

	// If all numbers have a corresponding pair, return -1
	return -1
}
