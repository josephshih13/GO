package main

import (
	"fmt"
	"strconv"
)

func append_range(ranges []string, first, last int) []string {
	combined_str := ""
	if first == last {
		combined_str = strconv.Itoa(first)
	} else {
		first_s, last_s := strconv.Itoa(first), strconv.Itoa(last)
		combined_str = first_s + "->" + last_s
	}

	return append(ranges, combined_str)
}

func summaryRanges(nums []int) []string {
	ret := []string{}
	first, last := -1, -1
	for i, v := range nums {
		if i == 0 {
			first, last = v, v
			continue
		}
		if v == last+1 {
			last++
		} else {
			ret = append_range(ret, first, last)
			first, last = v, v
		}
	}
	if len(nums) != 0 {
		ret = append_range(ret, first, last)
	}
	return ret
}

func main() {
	fmt.Println("123")
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{-1}))
	fmt.Println(summaryRanges([]int{0}))
}
