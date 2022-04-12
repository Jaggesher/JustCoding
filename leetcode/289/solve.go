package main

import "fmt"

func main() {
	fmt.Println("Case 1:")
	var arr [][]int = [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println(arr)
	gameOfLife(arr)
	fmt.Println(arr)
}

/***
 * Time: O(m*n)
 * Space: O(1)
 */
func gameOfLife(board [][]int) {
	var dirs [8][2]int = [8][2]int{{-1, -1}, {1, 1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, -1}, {-1, 1}}
	for i, row := range board {
		for j, vl := range row {
			neighbors := 0
			for _, dir := range dirs {
				u, v := i+dir[0], j+dir[1]
				if u < 0 || v < 0 || u >= len(board) || v >= len(row) {
					continue
				}
				neighbors += (board[u][v] % 2)
			}
			if vl == 1 && (neighbors < 2 || neighbors > 3) {
				board[i][j] = 3
			} else if vl == 0 && neighbors == 3 {
				board[i][j] = 2
			}
		}

	}

	for i, row := range board {
		for j, vl := range row {
			if vl == 2 {
				board[i][j] = 1
			} else if vl == 3 {
				board[i][j] = 0
			}
		}
	}
}
