package main

import "fmt"

func main() {
	var board [][]byte = [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println("Case 1:", board)
}

func solveSudoku(board [][]byte) {
	var row, column, grid [10][10]bool
	var emptyBox [][2]int = make([][2]int, 0, 90)
	for i, rowItem := range board {
		for j, vl := range rowItem {
			if vl != '.' {
				row[i][vl-'0'] = true
				column[j][vl-'0'] = true
				grid[((i/3)*3)+(j/3)][vl-'0'] = true
			} else {
				emptyBox = append(emptyBox, [2]int{i, j})
			}
		}
	}
	backtrack(0, emptyBox, board, &row, &column, &grid)
}

func backtrack(n int, emptyBox [][2]int, board [][]byte, row, column, grid *[10][10]bool) bool {
	if n == len(emptyBox) {
		return true
	}
	i, j, k := emptyBox[n][0], emptyBox[n][1], ((emptyBox[n][0]/3)*3)+(emptyBox[n][1]/3)
	for num := 1; num < 10; num++ {
		if !row[i][num] && !column[j][num] && !grid[k][num] {
			board[i][j] = byte(num + '0')
			row[i][num], column[j][num], grid[k][num] = true, true, true
			if backtrack(n+1, emptyBox, board, row, column, grid) {
				return true
			}
			board[i][j] = '.'
			row[i][num], column[j][num], grid[k][num] = false, false, false
		}
	}

	return false
}
