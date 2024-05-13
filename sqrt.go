package piscine

func Sqrt(nb int) int {
	result := 0
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i
			break
		}
	}
	return result
}
