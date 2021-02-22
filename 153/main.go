package main

import (
	"fmt"
)

func minint(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	m := (l + r) / 2
	for r-l > 1 {
		if nums[m] < nums[r] && nums[m] < nums[l] {
			r = m
		} else if nums[r] < nums[m] && nums[r] < nums[l] {
			l = m
		} else {
			r = m
		}
		m = (l + r) / 2
	}
	return minint(nums[l], minint(nums[r], nums[m]))
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
