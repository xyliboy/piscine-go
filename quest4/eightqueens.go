package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solver(0, &board)
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}

func isSafe(row, col int, board *[8]int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row || abs(r-row) == abs(c-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solver(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}
	for row := 0; row < 8; row++ {
		if isSafe(row, col, board) {
			board[col] = row
			solver(col+1, board)
		}
	}
}
