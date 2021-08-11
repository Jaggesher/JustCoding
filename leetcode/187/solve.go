package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
func findRepeatedDnaSequences(s string) []string {
	var ans []string = make([]string, 0)
	var flag map[string]int = make(map[string]int)
	for i := 10; i <= len(s); i++ {
		tm := s[(i - 10):i]
		tempStatus := flag[tm]
		if tempStatus == 1 {
			ans = append(ans, tm)
			flag[tm] = 2
		} else if tempStatus == 0 {
			flag[tm] = 1
		}
	}
	return ans
}
