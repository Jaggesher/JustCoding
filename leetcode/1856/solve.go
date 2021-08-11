package main

import "fmt"

func main() {
	//var arr []int = []int{9, 8, 7} //168
	//var arr []int = []int{2, 3, 4, 5, 1} //168
	//var arr []int = []int{9, 9, 9, 100} //10000
	//var arr []int = []int{2, 5, 4, 2, 4, 5, 3, 1, 2, 4} //50
	//var arr []int = []int{2, 3, 3, 1, 2} // 18
	//var arr []int = []int{3, 1, 5, 6, 4, 2} //60
	//var arr []int = []int{1, 2, 3, 2}       // 14
	//var arr []int = []int{1, 1, 3, 2, 2, 2, 1, 5, 1, 5} //25
	var arr []int = []int{10000000, 9999999} //25

	fmt.Println(maxSumMinProduct(arr))
}

func maxSumMinProduct(nums []int) int {
	const mod = 1000000000 + 7
	var ans uint64 = 0
	var sum []uint64 = make([]uint64, len(nums)+1)
	var history [][2]int = make([][2]int, 0)

	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + uint64(nums[i])
		flag := true
		for j := 0; j < len(history); j++ {
			if history[j][1] >= nums[i] {
				history[j][1] = nums[i]
				history = history[0 : j+1]
				flag = false
			}
		}
		if flag {
			history = append(history, [2]int{i, nums[i]})
		}
		for _, val := range history {
			ans = max(ans, (sum[i+1]-sum[val[0]])*uint64(val[1]))
		}
		//fmt.Println(i, history)
	}
	return int(ans % mod)
}

func max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
