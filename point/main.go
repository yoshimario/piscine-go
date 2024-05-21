package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func fixedPoint(num int) string {
	output := ""
	for ; num > 0; num = num / 10 {
		output = string(rune(num%10+'0')) + output
	}
	return output
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
	var print string = "x = " + fixedPoint(ptr.x) + ", y = " + fixedPoint(ptr.y) + "\n"
	for _, char := range print {
		z01.PrintRune(char)
	}
}

func main() {
	points := &point{}
	setPoint(points)
}
