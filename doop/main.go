package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}

	firstNum, firstNumOK := atoi(args[1])
	operator := args[2]
	secondNum, secondNumOK := atoi(args[3])

	if !firstNumOK || !secondNumOK {
		return
	}

	var result int64
	var overflow bool

	switch operator {
	case "+":
		result, overflow = add(firstNum, secondNum)
	case "-":
		result, overflow = subtract(firstNum, secondNum)
	case "*":
		result, overflow = multiply(firstNum, secondNum)
	case "/":
		if secondNum == 0 {
			printString("No division by 0\n")
			return
		}
		result = firstNum / secondNum
	case "%":
		if secondNum == 0 {
			printString("No modulo by 0\n")
			return
		}
		result = firstNum % secondNum
	default:
		return
	}

	if overflow {
		return
	}

	printInt(result)
	z01.PrintRune('\n')
}

func atoi(s string) (int64, bool) {
	var num int64
	sign := int64(1)
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int64(c-'0')
	}
	return sign * num, true
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func printInt(num int64) {
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for num > 0 {
		digits = append([]rune{rune(num%10) + '0'}, digits...)
		num = num / 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func add(a, b int64) (int64, bool) {
	result := a + b
	if (result > a) == (b > 0) {
		return result, false
	}
	return 0, true
}

func subtract(a, b int64) (int64, bool) {
	result := a - b
	if (result < a) == (b > 0) {
		return result, false
	}
	return 0, true
}

func multiply(a, b int64) (int64, bool) {
	result := a * b
	if b != 0 && result/b != a {
		return 0, true
	}
	return result, false
}
