package main

import "fmt"

func main() {
	fmt.Println("Case 1:", longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println("Case 2:", longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	fmt.Println("Case 3:", longestIncreasingPath([][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}))
}

type Box struct {
	x int
	y int
}

func longestIncreasingPath(matrix [][]int) int {
	var ans int = 1
	var tracker [201][201][3]int
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := make([]Box, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tracker[i][j][0], tracker[i][j][1], tracker[i][j][2] = 1, 1, 1
			queue = append(queue, Box{x: i, y: j})
		}
	}

	for len(queue) > 0 {
		temp := make([]Box, 0)
		for _, node := range queue {
			ans = max(ans, max(tracker[node.x][node.y][0], tracker[node.x][node.y][1]))
			tracker[node.x][node.y][2] = 0

			for _, dir := range dirs {
				x, y := node.x+dir[0], node.y+dir[1]
				if x < 0 || y < 0 || x >= m || y >= n {
					continue
				}
				if matrix[node.x][node.y] < matrix[x][y] && (tracker[node.x][node.y][0]+1) > tracker[x][y][0] {
					tracker[x][y][0] = tracker[node.x][node.y][0] + 1
					if tracker[x][y][2] == 0 {
						tracker[x][y][2] = 1
						temp = append(temp, Box{x: x, y: y})
					}
				} else if matrix[node.x][node.y] > matrix[x][y] && (tracker[node.x][node.y][1]+1) > tracker[x][y][1] {
					tracker[x][y][1] = tracker[node.x][node.y][1] + 1
					if tracker[x][y][2] == 0 {
						tracker[x][y][2] = 1
						temp = append(temp, Box{x: x, y: y})
					}
				}
			}
		}
		queue = temp
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
