package main

import (
	"fmt"
	"sort"
)

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	pre, l := nums[0], 1
	ret := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre+1 {
			l++
		}
		ret = maxint(l, ret)
		pre = nums[i]
	}
	return ret
}

func main() {
	fmt.Println("123")
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))

}
