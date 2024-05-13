package piscine

func FindNextPrime(nb int) int {
	for {
		if nb < 2 {
			nb = 2
		} else if nb == 2 || nb == 3 {
			return nb
		} else if nb%2 == 0 {
			nb++
			continue
		}

		isPrime := true
		for i := 3; i*i <= nb; i += 2 {
			if nb%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return nb
		}
		nb++
	}
}
