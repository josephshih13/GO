package main

import (
	"fmt"
)

func ttt(arr [3]int) {
	arr[0] = 100
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	ttt(arr)
	fmt.Println(arr)
}
