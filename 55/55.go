package main

import "fmt"

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	maxreach, l := 0, len(nums)
	for i := 0; i <= maxreach; i++ {
		maxreach = maxint(i+nums[i], maxreach)
		if maxreach >= l-1 {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
