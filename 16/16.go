package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// fmt.Println(nums)
	maxp, minp := math.MinInt32, math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// fmt.Println(i, nums)
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == target {
				// fmt.Println(l, r, i)
				return target
			}
			if sum < target {
				if sum > maxp {
					maxp = sum
				}
				l++
			} else {
				if sum < minp {
					minp = sum
				}
				r--
			}
		}
		// fmt.Println(i)
		// fmt.Println(maxp, minp)
	}
	// ans := 0
	if target-maxp > minp-target {
		return minp
	}
	return maxp
}

func main() {
	// sl := make([][][]int, 9, 9, 9)
	// fmt.Println(sl)
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
