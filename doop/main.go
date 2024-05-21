package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := []string{
		"doop",  // Simulating the command name, args[0]
		"1",     // args[1]
		"+",     // args[2]
		"1",     // args[3]
	}

	if len(args) != 4 {
		return
	}

	a, ok1 := atoi(args[1])
	operator := args[2]
	b, ok2 := atoi(args[3])

	if !ok1 || !ok2 {
		return
	}

	var result int64
	var overflow bool

	switch operator {
	case "+":
		result, overflow = add(a, b)
	case "-":
		result, overflow = subtract(a, b)
	case "*":
		result, overflow = multiply(a, b)
	case "/":
		if b == 0 {
			printString("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			printString("No modulo by 0\n")
			return
		}
		result = a % b
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
	var n int64
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
		n = n*10 + int64(c-'0')
	}
	return sign * n, true
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func printInt(n int64) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n = n / 10
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
