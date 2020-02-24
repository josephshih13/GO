package main

import (
	"fmt"
)

func dfs(now int, left int, s []byte, ans []string) []string {
	if now == 0 {
		for i := 0; i < left; i++ {
			s = append(s, ')')
		}
		ans = append(ans, string(s))
		return ans
	}
	if left > 0 {
		ss := append(s, ')')
		ans = dfs(now, left-1, ss, ans)
	}
	ss := append(s, '(')
	ans = dfs(now-1, left+1, ss, ans)
	return ans
}

func generateParenthesis(n int) []string {
	ans := []string{}
	if n == 0 {
		return ans
	}
	return dfs(n, 0, []byte{}, ans)
}

func main() {
	// s := "l12"
	// s[0] = 'f'
	// fmt.Println(s)
	// fmt.Println(strconv.Atoi(string(s[1])))
	fmt.Println(generateParenthesis(3))

}
