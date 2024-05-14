package piscine

import "github.com/01-edu/z01"

var (
	board     [8]int
	boardSize = 8
)

func EightQueens() {
	// Initialize the board
	for row := 0; row < boardSize; row++ {
		board[row] = -1
	}

	// Start placing queens
	row := 0
	for row >= 0 {
		board[row]++
		for board[row] < boardSize && !isSafe(row) {
			board[row]++
		}
		if board[row] < boardSize {
			if row == boardSize-1 {
				// Print the solution directly here
				for _, col := range board {
					// Print the digit by calculating its ASCII value
					n := col + 1
					if n == 0 {
						z01.PrintRune('0' + rune(n))
					} else {
						digits := []rune{}
						for n > 0 {
							digits = append([]rune{rune('0' + n%10)}, digits...)
							n /= 10
						}
						for _, d := range digits {
							z01.PrintRune(d)
						}
					}
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
