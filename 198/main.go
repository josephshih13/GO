package main

import (
	"fmt"
)

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 3 {
		ret := 0
		for _, v := range nums {
			ret = maxint(v, ret)
		}
		return ret
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]
	dp[2] = nums[0] + nums[2]
	ret := maxint(dp[2], dp[1])
	for i := 3; i < len(nums); i++ {
		dp[i] = maxint(dp[i-2], dp[i-3]) + nums[i]
		ret = maxint(ret, dp[i])
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
