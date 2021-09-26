package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Case 1: ", numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println("Case 2: ", numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	var seen map[string]bool = make(map[string]bool)
	for _, email := range emails {
		tm := strings.Split(email, "@")
		tm[0] = strings.Split(tm[0], "+")[0]
		tm[0] = strings.ReplaceAll(tm[0], ".", "")
		seen[fmt.Sprintf("%s@%s", tm[0], tm[1])] = true
	}
	return len(seen)
}
