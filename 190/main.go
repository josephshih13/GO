package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	bits := []uint32{}
	for num != 0 {
		bits = append(bits, num%2)
		num /= 2
	}
	for len(bits) < 32 {
		bits = append(bits, 0)
	}
	var ret uint32
	for i := 0; i < len(bits); i++ {
		ret *= 2
		ret += bits[i]
	}
	return ret
}

func main() {
	fmt.Println("123")
	s := "12345"
	fmt.Println(s[1:4] == "234")
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// [[1,2],[3,5],[6,7],[8,10],[12,16]]
// [4,8]
