package main

import "fmt"

func main() {
	//fmt.Println(largestMerge("cabaa", "bcaaa"))
	fmt.Println(largestMerge("abcabc", "abdcaba"))
	//fmt.Println(largestMerge("aa", "aab"))
	//fmt.Println(largestMerge("aaa", "aaa"))
	//fmt.Println(largestMerge("uuurruuuruuuuuuuuruuuuu", "urrrurrrrrrrruurrrurrrurrrrruu"))
	//fmt.Println(largestMerge("uuur", "ur"))
	//fmt.Println(largestMerge("jxddxdxddddxddddjdxdddddd", "ddddxddxdjdddxddjjddjdjdxdddj"))
	//fmt.Println(largestMerge("dxdy", "dxdz"))
}

func largestMerge(word1 string, word2 string) string {
	var arr1 []rune = []rune(word1)
	var arr2 []rune = []rune(word2)
	var ans []rune = make([]rune, 0)
	for len(arr1) > 0 && len(arr2) > 0 {
		//fmt.Println(string(arr1), string(arr2), string(ans))
		if arr1[0] > arr2[0] {
			ans = append(ans, arr1[0])
			arr1 = arr1[1:]
		} else if arr2[0] > arr1[0] {
			ans = append(ans, arr2[0])
			arr2 = arr2[1:]
		} else {
			cnt := 0
			ans = append(ans, arr1[0])
			for (cnt+1) < len(arr1) && (cnt+1) < len(arr2) && arr1[cnt+1] == arr2[cnt+1] && arr1[cnt] <= arr1[cnt+1] {
				cnt++
				ans = append(ans, arr1[cnt])
			}
			cnt++
			if cnt == len(arr1) {
				arr2 = arr2[cnt:]
			} else if cnt == len(arr2) {
				arr1 = arr1[cnt:]
			} else if arr1[cnt] > arr2[cnt] {
				arr1 = arr1[cnt:]
			} else if arr1[cnt] < arr2[cnt] {
				arr2 = arr2[cnt:]
			} else {
				flag := 1
				if len(arr1) < len(arr2) {
					flag = 2
				}
				tm := cnt
				for tm < len(arr1) && tm < len(arr2) {
					if arr1[tm] > arr2[tm] {
						flag = 1
						break
					} else if arr2[tm] > arr1[tm] {
						flag = 2
						break
					}
					tm++
				}
				if flag == 1 {
					arr1 = arr1[cnt:]
				} else {
					arr2 = arr2[cnt:]
				}

			}
		}
	}
	ans = append(ans, arr1...)
	ans = append(ans, arr2...)
	return string(ans)
}

/*
	F: "jxddxdxddddx"
	E: "jxddxdxddddxddxdjdddxddjjddjdjdxdddjddddxddddjdxdddddd"
	O: "jxddxdxddddxddxdjdddxddjjddjdjdxdddjddddxddddjdxdddddd"
*/
