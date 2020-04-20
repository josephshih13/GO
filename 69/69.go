package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	i := 1
	for i*i <= x {
		i++
	}
	return i - 1
}

func main() {

	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(5))
	fmt.Println(mySqrt(6))
}
