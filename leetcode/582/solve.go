package main

import "fmt"

func main() {
	fmt.Println("Case 1:", killProcess([]int{1, 3, 10, 5}, []int{3, 0, 5, 3}, 5))
	fmt.Println("Case 2:", killProcess([]int{1}, []int{0}, 1))
}

func killProcess(pid []int, ppid []int, kill int) []int {
	var ans []int = make([]int, 0)
	var tree map[int][]int = make(map[int][]int)
	for i, vl := range ppid {
		tree[vl] = append(tree[vl], pid[i])
	}
	queue := []int{kill}
	for len(queue) > 0 {
		temp := make([]int, 0)
		for _, vl := range queue {
			ans = append(ans, vl)
			temp = append(temp, tree[vl]...)
		}
		queue = temp
	}
	return ans
}
