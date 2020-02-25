package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	a := sort.SearchInts(nums, target)
	if a == len(nums) || nums[a] != target {
		return []int{-1, -1}
	}
	b := sort.SearchInts(nums, target+1)
	return []int{a, b - 1}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1, 1, 1, 1, 1, 1}, 1))
}
