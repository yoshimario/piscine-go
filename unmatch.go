package piscine

func Unmatch(numbers []int) int {
	countMap := make(map[int]int)

	// Count occurrences of each number
	for _, number := range numbers {
		countMap[number]++
	}

	// Find the number with an odd count
	for _, number := range numbers {
		if countMap[number]%2 != 0 {
			return number
		}
	}

	return -1
}
