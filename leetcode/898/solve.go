package main

import "fmt"

func main() {
	var arr []int = []int{0}
	fmt.Println(subarrayBitwiseORs(arr))
}

func subarrayBitwiseORs(arr []int) int {
	var ans int = 0
	var flag map[int]bool = make(map[int]bool)
	cache := []int{}
	for _, vl := range arr {
		if !flag[vl] {
			ans++
			flag[vl] = true
		}
		tm := make([]int, 0)
		tm = append(tm, vl)
		for i := 0; i < len(cache); i++ {
			cache[i] |= vl
			if !flag[cache[i]] {
				ans++
				flag[cache[i]] = true
			}
			if tm[len(tm)-1] != cache[i] {
				tm = append(tm, cache[i])
			}
		}
		cache = tm
	}
	return ans
}
