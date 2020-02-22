package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	z := make(map[rune]int)
	tag, ans := 0, 0
	for i, r := range s {
		pre, prs := z[r]
		if prs == true && pre+1 > tag {
			tag = pre + 1
		}
		z[r] = i
		if i-tag+1 > ans {
			ans = i - tag + 1
		}
	}
	return ans
}

func main() {
	//test("  -1233mfgmfm")
	// fmt.Println(myAtoi("words and 987"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
