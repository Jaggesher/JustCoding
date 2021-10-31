package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
}

// Approach One
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(nums []int, st, end int) {
	for st < end {
		nums[st], nums[end] = nums[end], nums[st]
		st++
		end--
	}
}

// =========================================
// Approach Two
func rotate1(nums []int, k int) {
	count, ln := 0, len(nums)
	for start := 0; count < ln; start++ {
		current, prev := start, nums[start]
		for ok := true; ok; ok = start != current {
			count++
			next := (current + k) % ln
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
		}
	}
}
