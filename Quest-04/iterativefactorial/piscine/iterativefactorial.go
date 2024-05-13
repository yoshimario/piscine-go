package piscine

func IterativeFactorial(nb int64) int64 {
    // Factorial of negative numbers or values greater than 12 will overflow int64
    if nb < 0 || nb > 12 {
        return 0
    }

    result := int64(1)
    for i := int64(1); i <= nb; i++ {
        result *= i
        // Check for overflow
        if result <= 0 {
            return 0
        }
    }

    return result
}
