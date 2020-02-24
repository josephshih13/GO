package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 6, 9}
	for i := 0; i < 12; i++ {
		fmt.Println(i, sort.SearchInts(arr, i))
	}
}
