package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ml := math.MaxInt32
	for _, str := range strs {
		if len(str) < ml {
			ml = len(str)
		}
	}
	ans := 0
	// fmt.Println(strs[0][0])
	for i := 0; i < ml; i++ {
		for _, str := range strs {
			if str[i] != strs[0][i] {
				return strs[0][:ans]
			}
		}
		ans++
	}
	return strs[0][:ans]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
}
