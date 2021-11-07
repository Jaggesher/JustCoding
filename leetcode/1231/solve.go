package main

import "fmt"

func main() {
	fmt.Println("Case 1:", maximizeSweetness([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println("Case 2:", maximizeSweetness([]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8))
	fmt.Println("Case 3:", maximizeSweetness([]int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2))
}

func maximizeSweetness(sweetness []int, k int) int {
	st, end := sweetness[0], 0
	for _, vl := range sweetness {
		end += vl
		if st > vl {
			st = vl
		}
	}
	end /= (k + 1)

	for st < end {
		mid := (st + end + 1) / 2
		cuts, tempSum := 0, 0
		for _, vl := range sweetness {
			tempSum += vl
			if tempSum >= mid {
				cuts++
				tempSum = 0
			}
		}
		if cuts >= (k + 1) {
			st = mid
		} else {
			end = mid - 1
		}
	}
	return end
}
