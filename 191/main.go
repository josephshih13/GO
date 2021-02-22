package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	ret := 0
	for num != 0 {
		if num%2 == 1 {
			ret += 1
		}
		num /= 2
	}
	return ret
}
func main() {
	fmt.Println("123")
	s := "12345"
	fmt.Println(s[1:4] == "234")
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// [[1,2],[3,5],[6,7],[8,10],[12,16]]
// [4,8]
