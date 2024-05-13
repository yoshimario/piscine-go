package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1 // Negative index will return -1
	}
	if index == 0 {
		return 0 // Base case: fibonacci(0) = 0
	}
	if index == 1 {
		return 1 // Base case: fibonacci(1) = 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
