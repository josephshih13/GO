package main

import (
	"fmt"
	"math"
)

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	sum, ret := math.MinInt32, math.MinInt32
	for _, v := range nums {
		sum = maxint(sum+v, v)
		ret = maxint(ret, sum)
	}
	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
