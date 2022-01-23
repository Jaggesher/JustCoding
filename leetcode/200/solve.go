package main

import "fmt"

func main() {
	fmt.Println("Case 1:", numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println("Case 2:", numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}

/***
 * Time: O(M*N)
 * Space: O(min(M, N))
 */
func numIslands(grid [][]byte) int {
	var ans int = 0
	var dirs [][]int = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ans++
				queue := make([][]int, 0)
				queue = append(queue, []int{i, j})
				grid[i][j] = '0'
				for len(queue) > 0 {
					node := queue[0]
					queue = queue[1:]
					for _, dir := range dirs {
						u, v := dir[0]+node[0], dir[1]+node[1]
						if u < 0 || v < 0 || u >= len(grid) || v >= len(grid[0]) || grid[u][v] != '1' {
							continue
						}
						grid[u][v] = '0'
						queue = append(queue, []int{u, v})
					}
				}
			}
		}
	}
	return ans
}
