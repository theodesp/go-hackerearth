package basic_programming

import "fmt"

func PrintNQueens() {
	var N int
	fmt.Scanf("%d", &N)

	board := makeBoard(N)

	result := NQueens(board, 0, N)
	if result {
		fmt.Println("YES")
		for i := 0; i < N; i += 1 {
			for j := 0; j < N; j += 1 {
				fmt.Printf("%d ", board[i][j])
			}
			fmt.Println()
		}
	} else {
		fmt.Println("NO")
	}

}

func makeBoard(n int) [][]int {
	board := make([][]int, n, n)
	for i := 0; i < n; i += 1 {
		board[i] = make([]int, n, n)
	}
	return board
}


func NQueens(board [][]int, col, n int) bool {
	if col >= n {
		return true
	}
	for i, _ := range board {
		if isSafe(board, i, col, n) {
			// Place current queen at cell (i,j)
			board[i][col] = 1
			if NQueens(board, col+1, n) {
				return true
			}
			// remove  current queen from (i,j)
			// Backtrack
			board[i][col] = 0
		}
	}
	// If the queen cannot be placed in any row in
	// this colum col  then return false
	return false
}

func isSafe(board [][]int, row, col, n int) bool {
	// Check this row on left side
	for i := 0; i < col; i += 1 {
		if board[row][i] == 1 {
			return false
		}
	}
	// Check upper diagonal on left side
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	// Check lower diagonal on left side
	for i, j := row, col; j >= 0 && i < n; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}
