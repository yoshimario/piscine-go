package piscine

func Abort(a, b, c, d, e int) int {
	// Step 1: Create a slice with the input values
	numbers := []int{a, b, c, d, e}

	// Step 2: Sort the first five elements manually using a simple sorting algorithm
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Step 3: Return the median, which is the middle element in the sorted list
	return numbers[2]
}
