package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	f1, f2, f3 := 1, 2, 3
	n -= 2
	for i := 0; i < n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
