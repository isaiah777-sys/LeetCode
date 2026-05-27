package leetcode

func solveNQueens(n int) [][]string {
	var results [][]string

	// Initialize an empty board
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// Track attacked lines
	cols := make(map[int]bool)
	diag1 := make(map[int]bool) // Row - Col constant
	diag2 := make(map[int]bool) // Row + Col constant

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			results = append(results, constructBoard(board))
			return
		}

		for col := 0; col < n; col++ {
			d1 := row - col
			d2 := row + col

			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			// Place queen
			board[row][col] = 'Q'
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			// Move to next row
			backtrack(row + 1)

			// Backtrack (Remove queen)
			board[row][col] = '.'
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	backtrack(0)
	return results
}

func constructBoard(board [][]byte) []string {
	var res []string
	for _, row := range board {
		res = append(res, string(row))
	}
	return res
}
