package main

import (
	"fmt"
)

func main() {
	board1 := [][]byte{
		{'B', 'B', 'B', '.', 'W', 'W', 'B', 'W'},
		{'B', 'B', '.', 'B', '.', 'B', 'B', 'B'},
		{'.', 'W', '.', '.', 'B', '.', 'B', 'W'},
		{'B', 'W', '.', 'W', 'B', '.', 'B', '.'},
		{'B', 'W', 'W', 'B', 'W', '.', 'B', 'B'},
		{'.', '.', 'W', '.', '.', 'W', '.', '.'},
		{'W', '.', 'W', 'B', '.', 'W', 'W', 'B'},
		{'B', 'B', 'W', 'W', 'B', 'W', '.', '.'}}

	fmt.Println(checkMove(board1, 5, 6, 'B'))
}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	board[rMove][cMove] = color
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	for _, vl := range dir {
		if findIfPossible(board, rMove, cMove, vl[0], vl[1]) {
			return true
		}
	}
	return false
}

func findIfPossible(board [][]byte, rMove int, cMove int, rNext, cNext int) bool {
	st, end, count, color, cChange := rMove+rNext, cMove+cNext, 1, board[rMove][cMove], 1
	for st >= 0 && st < 8 && end >= 0 && end < 8 && board[st][end] != '.' && cChange < 3 {
		count++
		if board[st][end] != color {
			cChange++
			color = board[st][end]
		}
		if color == board[rMove][cMove] {
			break
		}
		st += rNext
		end += cNext
	}
	if count >= 3 && cChange == 3 {
		return true
	}
	return false
}
