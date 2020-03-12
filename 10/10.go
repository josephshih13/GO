package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	// dp := [len(p) + 1][len(s) + 1]bool{}
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	for i, c1 := range p {

		if c1 == '*' {
			dp[i+1][0] = dp[i-1][0]
		}
		for j, c2 := range s {
			// if i == 8 && j == 7 {
			// 	fmt.Println(string(p[i-1]), string(s[j]))
			// }
			if c1 == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if c1 == '*' {
				if p[i-1] == '.' {
					dp[i+1][j+1] = dp[i-1][j+1] || dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i-1][j+1] || (dp[i+1][j] && p[i-1] == s[j])
				}
			} else {
				dp[i+1][j+1] = dp[i][j] && c1 == c2
			}
		}
	}
	// for i := range dp {
	// 	fmt.Println(dp[i])
	// }
	return dp[len(p)][len(s)]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
