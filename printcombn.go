package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(arr)
	data := make([]int, n)
	combinationUtil(arr, r, n, 0, data, 0)
}

func combinationUtil(arr []int, r, n, index int, data []int, i int) {
	if index == n {
		printCombination(data, n)
		return
	}
	if i >= r {
		return
	}
	data[index] = arr[i]
	combinationUtil(arr, r, n, index+1, data, i+1)
	combinationUtil(arr, r, n, index, data, i+1)
}

func printCombination(data []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(data[i] + '0'))
	}
	if !areLastCombination(data) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		z01.PrintRune('\n')
	}
}

func areLastCombination(data []int) bool {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 9-(len(data)-1-i) {
			return false
		}
	}
	return true
}
