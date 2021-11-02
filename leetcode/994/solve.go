package main

import "fmt"

func main() {
	fmt.Println("Case 1:", orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println("Case 2:", orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println("Case 3:", orangesRotting([][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}))
}

func orangesRotting(grid [][]int) int {
	var ans int = 0
	var dirs [][]int = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var queue [][2]int = make([][2]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	if len(queue) > 0 {
		ans = -1
	}
	for temp := make([][2]int, 0); len(queue) > 0; queue, temp = temp, make([][2]int, 0) {
		for _, node := range queue {
			for _, dir := range dirs {
				i, j := node[0]+dir[0], node[1]+dir[1]
				if (i >= 0) && (j >= 0) && (i < len(grid)) && (j < len(grid[i])) && (grid[i][j] == 1) {
					grid[i][j] = 2
					temp = append(temp, [2]int{i, j})
				}
			}
		}
		ans++
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}
