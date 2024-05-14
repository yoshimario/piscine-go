package piscine

import "github.com/01-edu/z01"

var (
	board     [8]int
	boardSize = 8
)

func EightQueens() {
	// Initialize the board with -1
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

// Helper function to print a single digit number using z01.PrintRune
func printDigit(n int) {
	// Convert the number to its corresponding rune
	z01.PrintRune('0' + rune(n))
}
