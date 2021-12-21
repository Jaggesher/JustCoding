package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", climbStairs(2)) //2
	fmt.Println("Case 2: ", climbStairs(3)) //3
	fmt.Println("Case 3: ", climbStairs(4)) //5
}

func climbStairs(n int) int {
	first, sec := 1, 1
	for i := 1; i < n; i++ {
		first, sec = sec, first+sec
	}
	return sec
}
