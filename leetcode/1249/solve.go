package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println("Case 2: ", minRemoveToMakeValid("a)b(c)d"))
	fmt.Println("Case 3: ", minRemoveToMakeValid("))(("))
}

/***
 * Time: O(n)
 * Space: O(n)
 */
func minRemoveToMakeValid(s string) string {
	var data []rune = []rune(s)
	var stack []int = make([]int, 0)

	for index, vl := range data {
		if vl == '(' {
			stack = append(stack, index)
		} else if vl == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				data[index] = '$'
			}
		}
	}

	for _, vl := range stack {
		data[vl] = '$'
	}

	n := 0
	for _, vl := range data {
		if vl != '$' {
			data[n] = vl
			n++
		}
	}
	return string(data[:n])
}
