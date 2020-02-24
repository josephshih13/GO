package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n, pre := 1, nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != pre {
			pre = nums[i]
			nums[i], nums[n] = nums[n], nums[i]
			n++
		}
	}
	return n
}

func main() {
	// s := "l12"
	// s[0] = 'f'
	// fmt.Println(s)
	// fmt.Println(strconv.Atoi(string(s[1])))
	test := []int{1, 1, 2}
	fmt.Println(removeDuplicates(test))
	fmt.Println(test)
	test = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(test))
	fmt.Println(test)
}
