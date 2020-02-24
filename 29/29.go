package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	t := 1
	if dividend < 0 {
		t = -t
		dividend = -dividend
	}
	if divisor < 0 {
		t = -t
		divisor = -divisor
	}
	// fmt.Println(t)
	ans := 0
	for dividend >= divisor {
		dividend -= divisor
		ans += 1
	}
	if t < 0 {
		ans = -ans
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return math.MaxInt32
	}
	return ans
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(0, -3))
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(divide(math.MinInt32, -1))
}
