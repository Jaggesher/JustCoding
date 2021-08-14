package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", climbStairs(2)) //2
	fmt.Println("Case 2: ", climbStairs(3)) //3
	fmt.Println("Case 3: ", climbStairs(4)) //5
}

func climbStairs(n int) int {
	var arr []int = make([]int, 47)
	arr[0], arr[1] = 1, 1
	for i := 0; i < n; i++ {
		arr[i+1] += arr[i]
		arr[i+2] += arr[i]
	}
	return arr[n-1]
}
