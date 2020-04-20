package main

import "fmt"

func stair(a int) int {
	if a == 0 {
		return 1
	}
	ret := 1
	for i := 1; i <= a; i++ {
		ret *= i
	}
	// fmt.Println(ret)
	return ret
}

func uniquePaths(m int, n int) int {
	m, n = m-1, n-1
	if m < n {
		m, n = n, m
	}
	ret := 1
	for i := m + n; i > m; i-- {
		ret *= i
	}
	return ret / stair(n)
}

func main() {

	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
}
