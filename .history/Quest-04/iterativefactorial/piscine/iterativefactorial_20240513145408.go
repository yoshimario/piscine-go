func main() {
	arg := 13 // Example argument
	// Convert int to int64
	arg64 := int64(arg)
	result := piscine.IterativeFactorial(arg64)
	fmt.Printf("IterativeFactorial(%d) = %d\n", arg, result)

	// Handling negative input case
	negativeArg := -8552095782491927294
	// Convert int to int64
	negativeArg64 := int64(negativeArg)
	negativeResult := piscine.IterativeFactorial(negativeArg64)
	fmt.Printf("IterativeFactorial(%d) = %d\n", negativeArg, negativeResult)
}
