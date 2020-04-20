package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1.0 / fastpow(x, -n)
	}
	return fastpow(x, n)
}

func fastpow(x float64, n int) float64 {
	var ans float64
	ans = 1.0
	for n != 0 {
		if n%2 == 1 {
			ans *= x

		}
		x *= x
		n /= 2
	}
	// fmt.Println(ans)
	return ans
}

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
	// fmt.Println(myPow())
}
