package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	ans := 0
	for len(s) > 0 {
		switch {
		case strings.HasPrefix(s, "M"):
			ans += 1000
			s = s[1:]
		case strings.HasPrefix(s, "CM"):
			ans += 900
			s = s[2:]
		case strings.HasPrefix(s, "D"):
			ans += 500
			s = s[1:]
		case strings.HasPrefix(s, "CD"):
			ans += 400
			s = s[2:]
		case strings.HasPrefix(s, "C"):
			ans += 100
			s = s[1:]
		case strings.HasPrefix(s, "XC"):
			ans += 90
			s = s[2:]
		case strings.HasPrefix(s, "L"):
			ans += 50
			s = s[1:]
		case strings.HasPrefix(s, "XL"):
			ans += 40
			s = s[2:]
		case strings.HasPrefix(s, "X"):
			ans += 10
			s = s[1:]
		case strings.HasPrefix(s, "IX"):
			ans += 9
			s = s[2:]
		case strings.HasPrefix(s, "V"):
			ans += 5
			s = s[1:]
		case strings.HasPrefix(s, "IV"):
			ans += 4
			s = s[2:]
		case strings.HasPrefix(s, "I"):
			ans++
			s = s[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
