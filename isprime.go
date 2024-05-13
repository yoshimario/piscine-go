package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false // 1 is not a prime number
	}
	if nb <= 3 {
		return true // 2 and 3 are prime numbers
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false // Numbers divisible by 2 or 3 are not prime
	}
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false // Numbers divisible by i or i+2 are not prime
		}
	}
	return true
}
