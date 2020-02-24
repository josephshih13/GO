package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	n := len(nums)
	for n > 0 && nums[n-1] == val {
		n--
	}
	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n--
		}
		for n >= 0 && nums[n-1] == val {
			n--
		}
	}
	return n
}

func main() {
	// s := "l12"
	// s[0] = 'f'
	// fmt.Println(s)
	// fmt.Println(strconv.Atoi(string(s[1])))
	test := []int{3, 2, 2, 3}
	fmt.Println(removeElement(test, 3))
	fmt.Println(test)
	test = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(test, 2))
	fmt.Println(test)
}
