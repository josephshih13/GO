package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}
	// inv_cnt := 0
	// for idx, n := range nums {
	// 	if n <= 0 || n > l {
	// 		nums[idx] = -1
	// 		inv_cnt++
	// 	}
	// }
	// l2 := l - inv_cnt
	// i2 := l2
	// for i2 < l && nums[i2] == -1 {
	// 	i2++
	// }
	// for idx := 0; idx < l2; idx++ {
	// 	if nums[idx] == -1 {
	// 		nums[idx], nums[i2] = nums[i2], nums[idx]
	// 		for i2 < l && nums[i2] == -1 {
	// 			i2++
	// 		}
	// 	}
	// }
	// fmt.Println(nums)

	// for idx := 0 ; idx < l2; idx++ {

	// }
	for idx := 0; idx < l; idx++ {
		for nums[idx] > 0 && nums[idx] <= l && nums[idx] != idx+1 {
			tmp := nums[idx]
			if nums[tmp-1] == tmp {
				nums[idx] = -1
			} else {
				nums[idx], nums[tmp-1] = nums[tmp-1], nums[idx]
			}
		}
	}
	fmt.Println(nums)

	for idx, n := range nums {
		if n != idx+1 {
			return idx + 1
		}
	}
	return l + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 10}))
}
