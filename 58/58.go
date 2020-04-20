package main

import "fmt"

func lengthOfLastWord(s string) int {
	ret := 0
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if !flag && s[i] == ' ' {
			continue
		}
		flag = true
		if s[i] == ' ' {
			return ret
		}
		ret++
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord("  "))
}
