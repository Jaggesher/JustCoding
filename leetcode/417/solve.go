package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
	fmt.Println("Case 2: ", pacificAtlantic([][]int{{2, 1}, {1, 2}}))
	fmt.Println("Case 3: ", pacificAtlantic([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}))
}

func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int = make([][]int, 0)
	var pacificOcean, atlanticOcean [205][205]bool
	queuePacific, queueAtlantic := make([][]int, 0), make([][]int, 0)

	for i := 0; i < len(heights); i++ {
		pacificOcean[i][0], atlanticOcean[i][len(heights[0])-1] = true, true
		queuePacific = append(queuePacific, []int{i, 0})
		queueAtlantic = append(queueAtlantic, []int{i, len(heights[i]) - 1})
	}

	for i := 0; i < len(heights[0]); i++ {
		pacificOcean[0][i], atlanticOcean[len(heights)-1][i] = true, true
		queuePacific = append(queuePacific, []int{0, i})
		queueAtlantic = append(queueAtlantic, []int{len(heights) - 1, i})
	}

	bfs(heights, queuePacific, &pacificOcean)
	bfs(heights, queueAtlantic, &atlanticOcean)

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if atlanticOcean[i][j] && pacificOcean[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

func bfs(heights, queue [][]int, tracker *[205][205]bool) {
	var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		temp := make([][]int, 0)
		for _, vl := range queue {
			for _, dir := range dirs {
				i, j := vl[0]+dir[0], vl[1]+dir[1]
				if i < 0 || j < 0 || i >= len(heights) || j >= len(heights[i]) || tracker[i][j] == true {
					continue
				}
				if heights[i][j] >= heights[vl[0]][vl[1]] {
					tracker[i][j] = true
					temp = append(temp, []int{i, j})
				}
			}
		}
		queue = temp
	}
}
