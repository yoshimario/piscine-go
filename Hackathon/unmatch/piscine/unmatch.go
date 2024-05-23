package piscine

func Unmatch(numbers []int) int {
	// Sort the numbers slice
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Check for the unmatching element
	for len(numbers) > 0 {
		if len(numbers) == 1 {
			return numbers[0]
		}
		if numbers[0] == numbers[1] {
			numbers = numbers[2:]
		} else {
			return numbers[0]
		}
	}

	return -1
}
