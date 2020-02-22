package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func test(str string) int {
	fmt.Println(strings.TrimLeft(str, " "))
	s := strings.TrimLeft(str, " ")
	if s == "" {
		return 0
	}
	k := 1
	if s[0] == '-' {
		k = -1
		s = s[1:]
	} else if s[0] == '+' {
		k = 1
		s = s[1:]
	}
	idx := 0
	for idx < len(s) {
		if !unicode.IsDigit(rune(s[idx])) {
			break
		}
		idx++
	}
	s = s[:idx]
	fmt.Println(s)
	fmt.Println(k)
	ss, _ := strconv.Atoi(s)
	fmt.Println(ss * k)
	ans := ss * k
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	} else if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}

func myAtoi(str string) int {
	return test(str)
}

func main() {
	//test("  -1233mfgmfm")
	fmt.Println(myAtoi("words and 987"))
}
