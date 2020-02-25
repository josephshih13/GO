package main

import (
	"fmt"
)

func dfs(candidates []int, now int, sum int, target int, sol [][]int, set []int) [][]int {
	if now == len(candidates) {
		return sol
	}
	sol = dfs(candidates, now+1, sum, target, sol, set)
	for sum <= target {
		set = append(set, candidates[now])
		sum += candidates[now]
		if sum == target {
			tmp := make([]int, len(set))
			copy(tmp, set)
			sol = append(sol, tmp)
			// fmt.Println(set)
			// fmt.Println(set)
			// fmt.Println(target)
			// sol = append(sol, set)
			// fmt.Println(sol)
		}
		sol = dfs(candidates, now+1, sum, target, sol, set)
		// fmt.Println(sol)
	}
	return sol
}

func combinationSum(candidates []int, target int) [][]int {
	tmp := [][]int{}
	tmp2 := []int{}
	return dfs(candidates, 0, 0, target, tmp, tmp2)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{3, 2, 6, 7}, 11))
}
