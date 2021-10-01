package main

import "fmt"

func main() {
	fmt.Println("Case 1:", addStrings("11", "123"))
	fmt.Println("Case 2:", addStrings("456", "77"))
	fmt.Println("Case 3:", addStrings("0", "0"))
	fmt.Println("Case 4:", addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	var ans []rune = make([]rune, 0)
	l1, l2 := len(num1)-1, len(num2)-1
	extra := 0
	for l1 >= 0 || l2 >= 0 {
		if l1 >= 0 {
			extra += int(num1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			extra += int(num2[l2] - '0')
			l2--
		}
		ans = append(ans, rune('0'+(extra%10)))
		extra /= 10
	}
	if extra > 0 {
		ans = append(ans, '1')
	}
	for st, end := 0, len(ans)-1; st < end; st, end = st+1, end-1 {
		ans[st], ans[end] = ans[end], ans[st]
	}
	return string(ans)
}
