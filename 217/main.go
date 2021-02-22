package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	check := map[int]int{}
	for _, v := range nums {
		_, prs := check[v]
		if prs {
			return true
		}
		check[v] = 1
	}
	return false
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
