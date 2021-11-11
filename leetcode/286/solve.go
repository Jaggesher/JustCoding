package main

import "fmt"

func main() {
	rooms1 := [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}
	wallsAndGates(rooms1)
	fmt.Println("Case 1:", rooms1)

	rooms2 := [][]int{{-1}}
	wallsAndGates(rooms2)
	fmt.Println("Case 2:", rooms2)

	rooms3 := [][]int{{2147483647}}
	wallsAndGates(rooms3)
	fmt.Println("Case 3:", rooms3)
}

type box struct {
	i      int
	j      int
	weight int
}

func wallsAndGates(rooms [][]int) {
	var queue []box = make([]box, 0)
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, box{i: i, j: j, weight: 0})
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) > 0 {
		temp := make([]box, 0)
		for _, vl := range queue {
			for _, dir := range dirs {
				i, j, weight := vl.i+dir[0], vl.j+dir[1], vl.weight+1
				if i < 0 || j < 0 || i >= len(rooms) || j >= len(rooms[0]) || rooms[i][j] == -1 || rooms[i][j] <= weight {
					continue
				}
				rooms[i][j] = weight
				temp = append(temp, box{i: i, j: j, weight: weight})
			}
		}
		queue = temp
	}
}
