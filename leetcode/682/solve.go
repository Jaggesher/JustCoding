package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Case 1: ", calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println("Case 2: ", calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	fmt.Println("Case 3: ", calPoints([]string{"1"}))
}

/***
 * Time: O(n)
 * Space: O(n)
 */
func calPoints(ops []string) int {
	var records []int64 = make([]int64, 0)
	for _, vl := range ops {
		switch vl {
		case "+":
			records = append(records, (records[len(records)-1] + records[len(records)-2]))
		case "D":
			records = append(records, records[len(records)-1]*2)
		case "C":
			records = records[:len(records)-1]
		default:
			num, _ := strconv.ParseInt(vl, 10, 64)
			records = append(records, num)
		}
	}
	ans := 0
	for _, vl := range records {
		ans += int(vl)
	}
	return ans
}
