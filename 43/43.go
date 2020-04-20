package main

import (
	"fmt"
	"strings"
)

func str2intslice(str string) []int {
	ret := make([]int, 300)
	for i, j := len(str)-1, 0; i >= 0; i, j = i-1, j+1 {
		ret[j] = int(str[i] - '0')
	}
	return ret
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	s1, s2 := str2intslice(num1), str2intslice(num2)
	sum := str2intslice("0")
	for i := 0; i < 120; i++ {
		for j := 0; j < 120; j++ {
			sum[i+j] += s1[i] * s2[j]
		}
	}
	for i := 0; i < 250; i++ {
		sum[i+1] += sum[i] / 10
		sum[i] %= 10
	}
	// fmt.Println(sum)
	var ans strings.Builder
	start := false
	for i := 250; i >= 0; i-- {
		if !start && sum[i] == 0 {
			continue
		}
		start = true
		ans.WriteRune(rune('0' + sum[i]))
	}
	return ans.String()
}

func main() {
	// fmt.Println(str2intslice("1234"))
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	// fmt.Println(multiply())
	// fmt.Println(multiply())
}
