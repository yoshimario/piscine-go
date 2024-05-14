package piscine

import "github.com/01-edu/z01"

var (
	board     [8]int
	boardSize = 8
)

func EightQueens() {
	var row int

	// Initialize the board with -1
	for row < boardSize {
		board[row] = -1
		row++
	}
	row = 0

	// Backtracking algorithm to place queens
	for row >= 0 {
		board[row]++
		for board[row] < boardSize && !isSafe(row) {
			board[row]++
		}
		if board[row] < boardSize {
			if row == boardSize-1 {
				printSolution()
			} else {
				row++
				board[row] = -1
			}
		} else {
			row--
		}
	}
}

func isSafe(row int) bool {
	col := board[row]
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]+i == col+row || board[i]-i == col-row {
			return false
		}
	}
	return true
}

func printSolution() {
	for _, col := range board {
		PrintNbr(col + 1)
	}
	z01.PrintRune('\n')
}

// Implementing PrintNbr function using z01.PrintRune
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	digits := []rune{}
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
