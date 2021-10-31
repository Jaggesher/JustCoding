package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
}

func moveZeroes(nums []int) {
	for slow, fast := 0, 0; slow < len(nums); slow, fast = slow+1, fast+1 {
		for fast < len(nums) && nums[fast] == 0 {
			fast++
		}
		if fast < len(nums) {
			nums[slow] = nums[fast]
		} else {
			nums[slow] = 0
		}
	}
}
