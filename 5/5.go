package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	var ans_r, ans_l int
	ans := 0
	for i, _ := range s {
		// fmt.Println(s[i])
		now := 1
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			now += 2
		}
		if now > ans {
			ans, ans_r, ans_l = now, r, l
		}
		// fmt.Println("@@")
		if i+1 < len(s) && s[i] == s[i+1] {
			now, l, r = 2, i-1, i+2
			for l >= 0 && r < len(s) && s[l] == s[r] {
				l--
				r++
				now += 2
			}
			if now > ans {
				ans, ans_r, ans_l = now, r, l
			}
		}
	}
	// fmt.Println("???")
	return string(s[ans_l+1 : ans_r])
}

func main() {
	//test("  -1233mfgmfm")
	// fmt.Println(myAtoi("words and 987"))
	fmt.Println(longestPalindrome("a"))
}
