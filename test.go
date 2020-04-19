package main

import "fmt"

func ttt(arr [3]int) {
	arr[0] = 100
}

func main() {
	m := make(map[string]int)
	m["123"] = 123
	fmt.Println(m["123"])
}
