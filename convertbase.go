package piscine

// ConvertBase converts a number from baseFrom to baseTo
func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr from baseFrom to base 10
	nbrBase10 := baseToDecimal(nbr, baseFrom)

	// Convert nbrBase10 to baseTo
	result := decimalToBase(nbrBase10, baseTo)
	return result
}

// baseToDecimal converts a number from a given base to decimal
func baseToDecimal(nbr, baseFrom string) int {
	base := len(baseFrom)
	nbrDecimal := 0
	for _, digit := range nbr {
		nbrDecimal *= base
		for i, d := range baseFrom {
			if d == digit {
				nbrDecimal += i
				break
			}
		}
	}
	return nbrDecimal
}

// decimalToBase converts a decimal number to a given base
func decimalToBase(nbr int, baseTo string) string {
	base := len(baseTo)
	var result string
	for nbr > 0 {
		remainder := nbr % base
		result = string(baseTo[remainder]) + result
		nbr /= base
	}
	return result
}
