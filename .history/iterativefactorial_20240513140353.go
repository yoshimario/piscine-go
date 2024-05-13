package piscine

func IterativeFactorial(nb int) int {
	// Factorial of negative numbers or values greater than 12 will overflow int
	if nb < 0 || nb > 12 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
		// Check for overflow
		if result <= 0 {
			return 0
		}
	}

	return result
}
