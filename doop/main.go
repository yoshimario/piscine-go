package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}

	firstNum, firstNumValid := convertToInt64(args[1])
	operator := args[2]
	secondNum, secondNumValid := convertToInt64(args[3])

	if !firstNumValid || !secondNumValid {
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
	printString("\n")
}

func convertToInt64(s string) (int64, bool) {
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
	os.Stdout.WriteString(s)
}

func printInt(num int64) {
	var digits []byte
	isNegative := num < 0
	if isNegative {
		num = -num
	}
	if num == 0 {
		digits = append(digits, '0')
	}
	for num > 0 {
		digits = append(digits, byte(num%10)+'0')
		num /= 10
	}
	if isNegative {
		digits = append(digits, '-')
	}
	for i := len(digits) - 1; i >= 0; i-- {
		os.Stdout.WriteByte(digits[i])
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
