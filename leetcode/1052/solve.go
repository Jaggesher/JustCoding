package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var ans int = 0
	mx, tempSum := 0, 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			ans += customers[i]
			customers[i] = 0
		}

		if (i - minutes) >= 0 {
			tempSum -= customers[i-minutes]
		}

		tempSum += customers[i]
		if mx < tempSum {
			mx = tempSum
		}
	}
	return ans + mx
}
