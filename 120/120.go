package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	//mi := 1<<60
	ma := make(map[[2]int]int)
	var trav func(int,int) int
	trav = func(i,j int)int{
		if i == len(triangle) -1 {
			return triangle[i][j]
		}
		tmp:=[2]int{i,j}
		_,ex := ma[tmp]
		if ex{
			return ma[tmp]
		}
		a := trav(i+1,j)
		b:= trav(i+1,j+1)
		if a < b{
			ma[tmp] = a + triangle[i][j]
		} else{
			ma[tmp] = b+triangle[i][j]
		}
		return ma[tmp]
	}
	return trav(0,0)
}

func main() {
	a:=[][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	fmt.Println(minimumTotal(a))
	aa:=[][]int{
		{-1},
		{2,3},
		{1,-1,-3},
	}
	fmt.Println(minimumTotal(aa))
}
