package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i-1] += digits[i] / 10
			digits[i] %= 10
		}
	}
	t := 0
	if digits[0] > 9 {
		t += digits[0] / 10
		digits[0] %= 10
		ret := make([]int, len(digits)+1)
		ret[0] = t
		copy(ret[1:], digits)
		return ret
	}
	return digits
}
func main() {
	// m := make(map[string]int)
	// m["123"] = 123
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
