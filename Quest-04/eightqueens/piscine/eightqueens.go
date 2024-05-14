package piscine

import "github.com/01-edu/z01"

var board [8]int
var boardSize = 8

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
				// Print the solution directly here
				for _, col := range board {
					printDigit(col + 1)
				}
				z01.PrintRune('\n')
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

// Helper function to print a single digit number
func printDigit(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		// Instead of using '0' directly, use its ASCII code value
		digits = append([]rune{rune(48 + n%10)}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
