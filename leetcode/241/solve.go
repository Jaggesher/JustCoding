package main

import (
	"fmt"
)

func main() {
	fmt.Println("Case 1:", diffWaysToCompute("2-1-1"))
	fmt.Println("Case 2:", diffWaysToCompute("2*3-4*5"))
}

func diffWaysToCompute(expression string) []int {
	var nums []int = make([]int, 0, 20)
	var operator []rune = make([]rune, 0, 20)
	tm := 0
	for _, vl := range expression {
		if vl == '+' || vl == '*' || vl == '-' {
			nums = append(nums, tm)
			tm = 0
			operator = append(operator, vl)
		} else {
			tm = (tm * 10) + int(vl-'0')
		}
	}
	nums = append(nums, tm)
	return solve(0, len(nums)-1, nums, operator)
}

func solve(st, end int, nums []int, operator []rune) []int {
	if st == end {
		return []int{nums[st]}
	}
	ans := make([]int, 0)
	for i := st; i < end; i++ {
		tm1, tm2 := solve(st, i, nums, operator), solve(i+1, end, nums, operator)
		for _, left := range tm1 {
			for _, right := range tm2 {
				ans = append(ans, calc(left, right, operator[i]))
			}
		}
	}
	return ans
}

func calc(num1, num2 int, operator rune) int {
	if operator == '+' {
		return num1 + num2
	}
	if operator == '-' {
		return num1 - num2
	}
	return num1 * num2
}
