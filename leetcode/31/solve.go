package main

import "fmt"

/***
 * ALGORITHM
 * Step 1: Find max "k" such that Array[k] < Array[k+1]
 * Step 2: Find max "l" such that Array[k] < Array[l]
 * Step 3: Swap between Array[k] & Array[l]
 * Step 4: Reverse the sequence from Array[k+1] to Array[n]
 * NB: If no such "k" exists then it's the highest possible permutation
 * Ref: https://en.wikipedia.org/wiki/Permutation
 */

func main() {
	arr1 := []int{1, 2, 3}
	nextPermutation(arr1)
	fmt.Println("Case 1:", arr1)

	arr2 := []int{3, 2, 1}
	nextPermutation(arr2)
	fmt.Println("Case 2:", arr2)
}

func nextPermutation(nums []int) {
	k, l, n := -1, -1, len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			k, l = i-1, i
			break
		}
	}

	if k != -1 {
		for j := n - 1; j > l; j-- {
			if nums[j] > nums[k] {
				l = j
			}
		}
		nums[k], nums[l] = nums[l], nums[k]
	}

	reverse(nums, k+1, n-1)
}

func reverse(arr []int, st, end int) {
	for st < end {
		arr[st], arr[end] = arr[end], arr[st]
		st++
		end--
	}
}
