package main

import "fmt"

type NumArray struct {
	nums []int
	bIT  []int
}

func Constructor(nums []int) NumArray {
	bIT := make([]int, len(nums)+1)
	arr := make([]int, len(nums))
	na := NumArray{nums: arr, bIT: bIT}

	for i, vl := range nums {
		na.Update(i, vl)
	}

	return na
}

func (this *NumArray) Update(index int, val int) {
	i, vl := index+1, val-this.nums[index]
	this.nums[index] = val
	for i < len(this.bIT) {
		this.bIT[i] += vl
		i += (i & (-i))
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.SumBIT(right+1) - this.SumBIT(left)
}

func (this *NumArray) SumBIT(index int) int {
	sum := 0
	for i := index; i > 0; i -= i & -i {
		sum += this.bIT[i]
	}
	return sum
}

func main() {
	tm := Constructor([]int{1, 3, 5})
	//fmt.Println(tm.bIT)
	fmt.Println(tm.SumRange(0, 2))
	tm.Update(1, 2)
	fmt.Println(tm.SumRange(0, 2))
}
