package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1:", subsetsWithDup([]int{1, 2, 2}))
	fmt.Println("Case 2:", subsetsWithDup([]int{0}))
	fmt.Println("Case 3:", subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

/***
 * Time: (n log n) + 2^n
 * Space: n+ 2^n
 */
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var flag map[string]bool = make(map[string]bool, 1024)
	var ans [][]int = make([][]int, 0, 1024)
	var baktrack func(n int, soFar []int)
	baktrack = func(n int, soFar []int) {
		if n >= len(nums) {
			key := ""
			for _, el := range soFar {
				key = fmt.Sprintf("%s#%d", key, el)
			}
			if _, ok := flag[key]; !ok {
				temp := make([]int, len(soFar))
				copy(temp, soFar)
				ans = append(ans, temp)
				flag[key] = true
			}
			return
		}
		baktrack(n+1, soFar)
		baktrack(n+1, append(soFar, nums[n]))
	}
	baktrack(0, []int{})
	return ans
}
