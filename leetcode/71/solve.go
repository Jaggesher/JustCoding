package main

import "fmt"

func main() {
	fmt.Println("Case 1: ", simplifyPath("/home/"))
	fmt.Println("Case 2: ", simplifyPath("/../"))
	fmt.Println("Case 3: ", simplifyPath("/home//foo/"))
	fmt.Println("Case 4: ", simplifyPath("/home/.erer/"))
	fmt.Println("Case 5: ", simplifyPath("/jkj"))
	fmt.Println("Case 6: ", simplifyPath("/a/./b/../../c/"))
}

/***
 * Time: O(n)
 * Space: O(n)
 */
func simplifyPath(path string) string {
	var stack []string = make([]string, 0)
	temp := make([]rune, 0)
	for _, vl := range path {
		if vl == '/' {
			stack = AddOrRemove(stack, string(temp))
			temp = temp[0:0]
		} else {
			temp = append(temp, vl)
		}
	}
	stack = AddOrRemove(stack, string(temp))

	if len(stack) == 0 {
		return "/"
	}

	ans := ""
	for _, vl := range stack {
		ans += "/" + vl
	}

	return ans
}

func AddOrRemove(stack []string, data string) []string {
	if data == ".." && len(stack) > 0 {
		return stack[:len(stack)-1]
	}

	if data != "." && data != "" && data != ".." {
		return append(stack, data)
	}
	return stack
}
