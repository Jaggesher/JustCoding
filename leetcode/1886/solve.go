package main

import "fmt"

func main() {
	fmt.Println("Case 1:", findRLEArray([][]int{{1, 3}, {2, 3}}, [][]int{{6, 3}, {3, 3}}))
	fmt.Println("Case 1:", findRLEArray([][]int{{1, 3}, {2, 1}, {3, 2}}, [][]int{{2, 3}, {3, 3}}))
}

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	var ans [][]int = make([][]int, 0)
	i1, i2, val, cnt := 0, 0, 0, 0
	for i1 < len(encoded1) && i2 < len(encoded2) {
		tm := encoded1[i1][0] * encoded2[i2][0]
		if tm != val {
			if cnt > 0 {
				ans = append(ans, []int{val, cnt})
			}
			val, cnt = tm, 1
		} else {
			cnt++
		}
		encoded1[i1][1]--
		encoded2[i2][1]--
		if encoded1[i1][1] <= 0 {
			i1++
		}
		if encoded2[i2][1] <= 0 {
			i2++
		}
	}
	if cnt > 0 {
		ans = append(ans, []int{val, cnt})
	}
	return ans
}
