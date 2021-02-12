package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	set := map[string]int{}
	for _, v := range wordDict {
		set[v] = 1
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s)+1; i++ {
		for j := i - 1; j >= 0; j-- {
			_, prs := set[s[j:i]]
			if dp[j] == 1 && prs {
				dp[i] = 1
				break
			}
		}
	}
	return dp[len(s)] == 1
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
