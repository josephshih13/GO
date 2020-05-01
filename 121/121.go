package main

import "fmt"

func maxProfit(prices []int) int {
	pre_min := 1<< 60
	ans := 0
	for _,v:=range prices{
		if v - pre_min > ans{
			ans = v - pre_min
		}
		if pre_min > v{
			pre_min = v
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
