package piscine

func SortIntegerTable(table []int) {
	length := len(table)
	for currentIndex := 0; currentIndex < length-1; currentIndex++ {
		for innerIndeex := 0; innerIndeex < length-currentIndex-1; innerIndeex++ {
			if table[innerIndeex] > table[innerIndeex+1] {
				table[innerIndeex], table[innerIndeex+1] = table[innerIndeex+1], table[innerIndeex]
			}
		}
	}
}
