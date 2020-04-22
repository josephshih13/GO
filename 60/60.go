package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dfs(now, n int, used []bool, ans [][]int, nums, arr []int) [][]int {
	if now == n {
		tmp := make([]int, n)
		copy(tmp, arr)
		ans = append(ans, tmp)
		return ans
	}
	for i, v := range used {
		if !v {
			used[i] = true
			arr[now] = nums[i]
			ans = dfs(now+1, n, used, ans, nums, arr)
			used[i] = false
		}
	}
	return ans
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	used := make([]bool, len(nums))
	arr := make([]int, len(nums))
	return dfs(0, len(nums), used, ans, nums, arr)
}

func getPermutation(n int, k int) string {
	cand := []int{}
	for i := 1; i <= n; i++ {
		cand = append(cand, i)
	}
	tmp := permute(cand)
	// fmt.Println(tmp)
	cc := tmp[k-1]
	var sb strings.Builder
	for _, v := range cc {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func main() {

	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
}
