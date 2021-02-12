package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}

	}

	for i := 0; i < l; i++ {
		fmt.Println(matrix[i])
	}

}

func main() {
	fmt.Println("123")
	rotate([][]int{{1, 2}, {3, 4}})

}
