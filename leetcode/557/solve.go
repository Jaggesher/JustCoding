package main

import "fmt"

func main() {
	fmt.Println("Case 1:", reverseWords("Let's take LeetCode contest"))
	fmt.Println("Case 2:", reverseWords("God Ding"))
}

func reverseWords(s string) string {
	arr := []rune(s)
	for st := 0; st < len(arr); st++ {
		if arr[st] == ' ' {
			continue
		}

		end := st
		for (end+1) < len(arr) && arr[end+1] != ' ' {
			end++
		}

		temp := end
		for st < end {
			arr[st], arr[end] = arr[end], arr[st]
			st, end = st+1, end-1
		}

		st = temp
	}
	return string(arr)
}
