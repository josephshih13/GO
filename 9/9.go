package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	// i := 0
	for i, j := 0, len(s)-1; i < len(s); i++ {
		fmt.Println(i, j)
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}
