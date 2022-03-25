package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
	fmt.Println("Case 2: ", twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}))
	fmt.Println("Case 3: ", twoCitySchedCost([][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}))
}

/***
 * Approach: Greedy
 * Time: O(n Log n) + O(n) => O(n Log n)
 * Space: O(n) // for sorting
 */
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(a, b int) bool {
		return Abs(costs[a][0]-costs[a][1]) > Abs(costs[b][0]-costs[b][1])
	})

	a, b, ans, n := 0, 0, 0, len(costs)/2

	for _, cost := range costs {
		if (cost[0] <= cost[1] && a < n) || b >= n {
			a, ans = a+1, ans+cost[0]
		} else {
			b, ans = b+1, ans+cost[1]
		}
	}

	return ans
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
