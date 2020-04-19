package main

import (
	"fmt"
)

func dfs(now, n int, used []bool, ans [][]int, nums, arr []int) [][]int {
	if now == n {
		tmp := make([]int, n)
		copy(tmp, arr)
		ans = append(ans, tmp)
		return ans
	}
	used_tmp := make(map[int]bool)
	for _, v := range nums {
		used_tmp[v] = false
	}
	for i, v := range used {
		if !v && !used_tmp[nums[i]] {
			used_tmp[nums[i]] = true
			used[i] = true
			arr[now] = nums[i]
			ans = dfs(now+1, n, used, ans, nums, arr)
			used[i] = false
		}
	}
	return ans
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	ans := [][]int{}
	used := make([]bool, len(nums))
	arr := make([]int, len(nums))
	return dfs(0, len(nums), used, ans, nums, arr)
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
