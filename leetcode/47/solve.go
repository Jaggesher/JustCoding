package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 3}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{6, 7, 8, 9, 1, 2}))
	fmt.Println(trap([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}))
}

func trap(height []int) int {
	var ans int = 0
	var monoQueue [][2]int = make([][2]int, 0)
	for _, vl := range height {
		if len(monoQueue) > 0 {
			if monoQueue[len(monoQueue)-1][0] > vl {
				monoQueue = append(monoQueue, [2]int{vl, 1})
			} else {
				index := 0
				for monoQueue[index][0] > vl {
					index++
				}
				temp := vl
				if index == 0 {
					temp = monoQueue[0][0]
				}
				cnt := 1
				for i := index; i < len(monoQueue); i++ {
					item := monoQueue[i]
					cnt += item[1]
					ans += ((temp - item[0]) * item[1])
				}
				monoQueue = monoQueue[:index]
				monoQueue = append(monoQueue, [2]int{vl, cnt})
			}
		} else if vl > 0 {
			monoQueue = append(monoQueue, [2]int{vl, 1})
		}
	}
	return ans
}

func trap2(height []int) int {
	var ans int = 0
	for _, vl := range height {
		fmt.Println(vl)
	}
	return ans
}
