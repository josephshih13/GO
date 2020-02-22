package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	ans := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		if area > ans {
			ans = area
		}
		if height[i] > height[j] {
			j--
		} else if height[i] < height[j] {
			i++
		} else {
			j--
			i++
		}

	}
	return ans
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
