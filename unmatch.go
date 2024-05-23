package piscine

func Unmatch(a []int) int {
	// Create a map to store the frequency of each element
	freqMap := make(map[int]int)

	// Count the frequency of each element in the slice
	for _, num := range a {
		freqMap[num]++
	}

	// Traverse the list in reverse to get the last unpaired element
	for i := len(a) - 1; i >= 0; i-- {
		if freqMap[a[i]]%2 != 0 {
			return a[i]
		}
	}

	// If all numbers have a corresponding pair, return -1
	return -1
}
