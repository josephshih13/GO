package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ans strings.Builder
	r := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		if i%numRows == 0 || i%numRows == numRows-1 {
			for j := i; j < len(s); j += r {
				ans.WriteByte(s[j])
			}
		} else {
			g := (numRows - 1 - i) * 2
			g = r - g
			for j := i; j < len(s); j += g {
				ans.WriteByte(s[j])
				g = r - g
			}
		}
	}
	return ans.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
