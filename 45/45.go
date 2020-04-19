package main

import (
	"fmt"
)

func jump(nums []int) int {

	now, l := 0, len(nums)
	if l <= 1 {
		return 0
	}
	cnt := 0
	for now < l-1 {
		if now+nums[now] >= l-1 {
			return cnt + 1
		}
		next_best, ni := -1, -1
		for jl := 1; jl <= nums[now]; jl++ {
			next_idx := now + jl
			if nums[next_idx]+next_idx > next_best {
				ni = next_idx
				next_best = nums[next_idx] + next_idx
			}
		}
		// fmt.Println(next_best, ni)
		now = ni
		cnt++
		// fmt.Println(now)
	}
	return -1
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1, 2, 3}))
}
