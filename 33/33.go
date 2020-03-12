package main

import (
	"fmt"
)

func is(nums []int, target int, l int, r int) int {
	m := (l + r) / 2
	if r-l <= 1 {
		if nums[l] == target {
			return l
		}
		return -1
	}
	ll, rr, mm := nums[l], nums[r], nums[m]
	if mm == target {
		return m
	}
	if mm > ll && mm > rr {
		if target > mm || target < rr {
			return is(nums, target, m, r)
		}
		return is(nums, target, l, m)
	} else if mm < ll && mm < rr {
		if target < mm || target > ll {
			return is(nums, target, l, m)
		}
		return is(nums, target, m, r)
	} else {
		if target > mm {
			return is(nums, target, m, r)
		}
		return is(nums, target, l, m)
	}
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	return is(nums, target, 0, len(nums)-1)
}

func main() {
	for i := 0; i < 8; i++ {
		fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, i))
	}
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
