package piscine

import "fmt"

const boardSize = 8

func EightQueens() {
	board := make([]int, boardSize)
	placeQueens(board, 0)
}

func placeQueens(board []int, row int) {
	if row == boardSize {
		printSolution(board)
		return
	}

	for col := 0; col < boardSize; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			placeQueens(board, row+1)
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]+i == col+row || board[i]-i == col-row {
			return false
		}
	}
	return true
}

func printSolution(board []int) {
	for _, col := range board {
		fmt.Print(col + 1)
	}
	fmt.Println()
}
