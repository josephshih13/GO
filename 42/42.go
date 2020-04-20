package main

import (
	"fmt"
)

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	l := len(height)
	dp := make([]int, l)
	m := -1
	for i := 0; i < l; i++ {
		if height[i] > m {
			m = height[i]
			dp[i] = m - height[i]
		} else {
			dp[i] = m - height[i]
		}
	}
	// fmt.Println(dp)
	m = -1
	for i := l - 1; i >= 0; i-- {
		if height[i] > m {
			m = height[i]
			dp[i] = minint(m-height[i], dp[i])
		} else {
			dp[i] = minint(m-height[i], dp[i])
		}
	}
	sum := 0
	for i := 0; i < l; i++ {
		sum += dp[i]
	}
	// fmt.Println(dp)
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

}
