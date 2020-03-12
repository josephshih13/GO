package main

import (
	"fmt"
)

func ttt(arr [3]int) {
	arr[0] = 100
}

func main() {
	sl := make([][][]int, 9, 9, 9)
	fmt.Println(sl)
}
