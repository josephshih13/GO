package main

import (
	"fmt"
)

func cross(inter1, inter2 []int) bool {
	// fmt.Println(inter1, inter2)
	x1, y1 := inter1[0], inter1[1]
	x2, y2 := inter2[0], inter2[1]
	if y1 < x2 || y2 < x1 {
		return false
	}
	return true
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{}
	x, y := newInterval[0], newInterval[1]
	for i := 0; i < len(intervals); i++ {
		if !cross(newInterval, intervals[i]) {
			ret = append(ret, intervals[i])
		} else {
			x = minint(x, intervals[i][0])
			y = maxint(y, intervals[i][1])
		}
	}

	// fmt.Println(x, y, ret)

	inter := []int{x, y}
	idx := -1
	for ii, v := range ret {
		if v[0] > x {
			idx = ii
			break
		}
	}

	if idx == -1 {
		idx = len(ret)
	}

	ret = append(ret, []int{})
	copy(ret[idx+1:], ret[idx:])
	ret[idx] = inter

	return ret

}

func main() {
	fmt.Println("123")
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))

}

// [[1,2],[3,5],[6,7],[8,10],[12,16]]
// [4,8]
