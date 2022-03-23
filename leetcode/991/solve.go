package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", brokenCalc(2, 3))
	fmt.Println("Case 2: ", brokenCalc(5, 8))
	fmt.Println("Case 3: ", brokenCalc(3, 10))
}

/***
 * Approach: Greedy
 * Time: O(log target)
 * Space: O(1)
 */

func brokenCalc(startValue int, target int) int {
	var ans int = 0
	for target > startValue {
		ans++
		if (target & 1) == 1 {
			target++
		} else {
			target = target >> 1
		}
	}
	return ans + startValue - target
}
