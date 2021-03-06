package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
	fmt.Println(minOperations([]int{3}, 3))
}

/***
 * Time: O(n)
 * Space: O(1)
 */
func minOperations(nums []int, x int) int {
	var st, end, sum, ans int = -1, len(nums) - 1, 0, -1
	for (st+1) < len(nums) && (sum+nums[st+1]) <= x {
		st++
		sum += nums[st]
		if sum == x {
			ans = st + 1
			break
		}
	}

	for st < end && end >= 0 {
		sum += nums[end]
		for sum > x && st >= 0 {
			sum -= nums[st]
			st--
		}

		if sum == x {
			tm := (st + 1) + (len(nums) - end)
			if ans == -1 || ans > tm {
				ans = tm
			}
		}

		end--
	}

	return ans
}

func minOperationsOld(nums []int, x int) int {
	var ans int = -1
	var tracker map[int]int = make(map[int]int)
	tracker[0] = 0
	sum := 0
	for i, vl := range nums {
		sum += vl
		if sum > x {
			break
		}
		tracker[sum] = i + 1
	}

	if _, ok := tracker[x]; ok {
		ans = tracker[x]
	}

	sum = x
	for i := len(nums) - 1; i >= 0; i-- {
		sum -= nums[i]
		if sum < 0 {
			break
		}

		index, ok := tracker[sum]
		if ok && (index < (i + 1)) {
			tm := index + (len(nums) - i)
			if ans == -1 || ans > tm {
				ans = tm
			}
		}
	}
	return ans
}
