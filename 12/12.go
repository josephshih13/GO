package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var str strings.Builder

	for num > 0 {
		switch {
		case num >= 1000:
			str.WriteString("M")
			num -= 1000
		case num >= 900:
			str.WriteString("CM")
			num -= 900
		case num >= 500:
			str.WriteString("D")
			num -= 500
		case num >= 400:
			str.WriteString("CD")
			num -= 400
		case num >= 100:
			str.WriteString("C")
			num -= 100
		case num >= 90:
			str.WriteString("XC")
			num -= 90
		case num >= 50:
			str.WriteString("L")
			num -= 50
		case num >= 40:
			str.WriteString("XL")
			num -= 40
		case num >= 10:
			str.WriteString("X")
			num -= 10
		case num >= 9:
			str.WriteString("IX")
			num -= 9
		case num >= 5:
			str.WriteString("V")
			num -= 5
		case num >= 4:
			str.WriteString("IV")
			num -= 4
		case num >= 1:
			str.WriteString("I")
			num -= 1
		}
		// fmt.Println(num)
	}

	return str.String()
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
