package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	var sb strings.Builder
	po := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i] - '0')
		}
		if j >= 0 {
			y = int(b[j] - '0')
		}
		d := (x + y + po) % 2
		po = (x + y + po) / 2
		// fmt.Println(po, d)
		sb.WriteRune(rune(d + '0'))
	}
	if po != 0 {
		sb.WriteRune(rune('1'))
	}
	var sb2 strings.Builder
	ss := sb.String()
	for i := len(ss) - 1; i >= 0; i-- {
		sb2.WriteByte(ss[i])
	}
	return sb2.String()
}

func main() {
	// m := make(map[string]int)
	// m["123"] = 123
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
