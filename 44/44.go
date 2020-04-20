package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	if s == "" && p == "*" {
		return true
	}
	if s == "" {
		return false
	}

	dp := [][]bool{}
	for idx := 0; idx < len(p)+1; idx++ {
		tmp := make([]bool, len(s)+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 0; i < len(p) && p[i] == '*'; i++ {
		dp[i+1][0] = true
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if p[i] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[i] == '*' {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i][j]
			} else {
				dp[i+1][j+1] = p[i] == s[j] && dp[i][j]
			}

		}
	}
	// fmt.Println(dp)
	return dp[len(p)][len(s)]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("abcde", "a*e"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("ab", "?*"))
}
