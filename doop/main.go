package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	operator := os.Args[2]
	b, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
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
			fmt.Println("No division by 0")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = a % b
	default:
		return
	}

	if overflow {
		return
	}

	fmt.Println(result)
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
