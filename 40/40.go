package main

import (
	"fmt"
)

func dfs(candidates []int, cnt []int, now int, sum int, target int, sol [][]int, set []int) [][]int {
	if now == len(candidates) {
		return sol
	}
	sol = dfs(candidates, cnt, now+1, sum, target, sol, set)
	for i := 0; i < cnt[now] && sum < target; i++ {
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
		sol = dfs(candidates, cnt, now+1, sum, target, sol, set)
		// fmt.Println(sol)
	}
	return sol
}

func combinationSum2(candidates []int, target int) [][]int {
	m := map[int]int{}
	for _, num := range candidates {
		_, e := m[num]
		if !e {
			m[num] = 1
		} else {
			m[num]++
		}
	}
	cnt := []int{}
	cand := []int{}
	for k, v := range m {
		cnt = append(cnt, v)
		cand = append(cand, k)
	}
	tmp := [][]int{}
	tmp2 := []int{}
	return dfs(cand, cnt, 0, 0, target, tmp, tmp2)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
