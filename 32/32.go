package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	st := []int{-1}
	// pre := -1
	ans := 0
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = append(st, i)
			} else if i-st[len(st)-1] > ans {
				ans = i - st[len(st)-1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}
