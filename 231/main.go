package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for ; n != 1; n = n >> 1 {
		if n%2 == 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("123")
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(15))
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(1024))
}
