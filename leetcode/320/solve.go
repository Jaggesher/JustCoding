package main

import "fmt"

func main() {
	fmt.Println("Case 1:", generateAbbreviations("word"))
	fmt.Println("Case 1:", generateAbbreviations("a"))
}

/***
 * Time: O(N * 2^N) => N
 * Memory: O(N)
 */
func generateAbbreviations(word string) []string {
	var ans []string = make([]string, 0, 1<<len(word))
	var backtrack func(int, int, string)

	backtrack = func(index, carry int, soFar string) {
		if index >= len(word) {
			if carry > 0 {
				soFar = fmt.Sprintf("%s%d", soFar, carry)
			}
			ans = append(ans, soFar)
			return
		}
		backtrack(index+1, carry+1, soFar)
		if carry > 0 {
			soFar = fmt.Sprintf("%s%d", soFar, carry)
		}
		backtrack(index+1, 0, fmt.Sprintf("%s%c", soFar, word[index]))
	}
	backtrack(0, 0, "")

	return ans
}
