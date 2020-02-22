package main

import (
	"fmt"
	"strconv"
)

func dfs(pos int, digits string, now []byte, ans []string) []string {
	phone := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if pos == len(digits) {
		if pos == 0 {
			return ans
		}
		fmt.Println(string(now))
		ans = append(ans, string(now))
		return ans
	}
	idx, _ := strconv.Atoi(string(digits[pos]))
	for _, c := range phone[idx] {
		fmt.Println(string(c))
		now[pos] = byte(c)
		ans = dfs(pos+1, digits, now, ans)
	}
	return ans
}

func letterCombinations(digits string) []string {
	ans := []string{}
	emp := make([]byte, len(digits))
	return dfs(0, digits, emp, ans)
}

func main() {
	// s := "l12"
	// s[0] = 'f'
	// fmt.Println(s)
	// fmt.Println(strconv.Atoi(string(s[1])))
	fmt.Println(letterCombinations(""))
}
