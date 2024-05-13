package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 12 {
		return 0 // Error: Values less than 0 or greater than 12 will result in overflow
	}
	if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
