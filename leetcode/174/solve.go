package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Case 1:", calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	fmt.Println("Case 2:", calculateMinimumHP([][]int{{0}}))
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			tm := 0
			if (i+1) < m && (j+1) < n {
				tm = int(math.Max(float64(dungeon[i+1][j]), float64(dungeon[i][j+1])))
			} else if (i + 1) < m {
				tm = dungeon[i+1][j]
			} else if (j + 1) < n {
				tm = dungeon[i][j+1]
			}
			if tm < 0 {
				dungeon[i][j] += tm
			}
		}
	}

	if dungeon[0][0] < 0 {
		return -dungeon[0][0] + 1
	}
	return 1
}
