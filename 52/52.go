package main

import (
	"fmt"
	"strings"
)

func validq(x, y int, queen [][]int) bool {
	for i := range queen {
		if queen[i][0] == x || queen[i][1] == y {
			return false
		}
		dx, dy := x-queen[i][0], y-queen[i][1]
		if dx == dy || dx == (-dy) {
			return false
		}
	}
	return true
}

func conv2ans(queens [][]int, n int) []string {
	ret := []string{}
	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := 0; j < n; j++ {
			if j != queens[i][1] {
				sb.WriteRune('.')
			} else {
				sb.WriteRune('Q')
			}
		}
		ret = append(ret, sb.String())
	}
	return ret
}

func dfs(now_x, n int, queens [][]int, ans [][]string) [][]string {
	if now_x == n {
		ans = append(ans, conv2ans(queens, n))
		return ans
	}
	for y := 0; y < n; y++ {
		if validq(now_x, y, queens) {
			queens = append(queens, []int{now_x, y})
			ans = dfs(now_x+1, n, queens, ans)
			queens = queens[:len(queens)-1]
		}
	}
	return ans
}

func totalNQueens(n int) int {
	ans := [][]string{}
	queens := [][]int{}
	ans = dfs(0, n, queens, ans)
	return len(ans)
}

func main() {
	fmt.Println(totalNQueens(4))
}
