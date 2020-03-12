package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			reverse(nums[i+1:])
			return
		}
	}
	reverse(nums)
}

func main() {
	t := []int{1, 3, 2}
	nextPermutation(t)
	fmt.Println(t)
	t = []int{3, 2, 1}
	nextPermutation(t)
	fmt.Println(t)
	t = []int{1, 1, 5}
	nextPermutation(t)
	fmt.Println(t)
	// fmt.Println(nextPermutation([]int{3, 2, 1}))
	// fmt.Println(nextPermutation([]int{1, 1, 5}))
}
