package main

import "fmt"

func setZeroes(matrix [][]int)  {
	if len(matrix)==0{
		return
	}
	zr,zc := []int{},[]int{}
	m,n:=len(matrix),len(matrix[0])
	if n==0{
		return
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==0{
				zr = append(zr,i)
				zc = append(zc,j)
			}
		}
	}
	for _,r:=range zr{
		for j:=0;j<n;j++{
			matrix[r][j]=0
		}
	}
	for _,c:= range zc{
		for i:=0;i<m;i++{
			matrix[i][c]=0
		}
	}
	return
}

func main() {
	a := [][]int{
		{1,1,1},
		{1,0,1},
		{1,1,1},
	}
	setZeroes(a)
	fmt.Println(a)
}
