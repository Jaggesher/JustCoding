package main

import "fmt"

func main() {
	fmt.Println("No test case")
}

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	temp := make([]int, len(nums)+1)
	for i, vl := range nums {
		temp[i+1] = temp[i] + vl
	}
	return NumArray{preSum: temp}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
