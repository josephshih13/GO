package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var b, p int64
	b, p = 0, 1
	a := int64(x)
	if a < 0 {
		p = -1
		a = -a
	}
	for a > 0 {
		b *= 10
		b += a % 10
		a /= 10
	}
	b *= p
	if b > math.MaxInt32 || b < math.MinInt32 {
		b = 0
	}
	return int(b)
}

func main() {
	fmt.Println(reverse(-100))
}
