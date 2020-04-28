package main

import (
	"fmt"
	"sort"
)

func mysearch(a []int, tar int) bool{
	tmp := sort.SearchInts(a,tar)
	if tmp == len(a) || a[tmp] != tar{
		return false
	}
	return true
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}
	//l,r := 0,len(matrix)-1
	//if target < matrix[l][0]{
	//	return false
	//}
	//if target > matrix[r][0]{
	//	return mysearch(matrix[r],target)
	//}
	//m := (l+r)/2
	arr := []int{}
	for i:=0;i<len(matrix);i++{
		arr = append(arr,matrix[i][0])
	}
	//fmt.Println(arr)
	ii := sort.SearchInts(arr,target)
	if ii<len(matrix)&&matrix[ii][0] == target{
		return true
	}
	ii--
	if ii <0{
		return false
	}
	return mysearch(matrix[ii],target)
}

func main() {
	a:=[][]int{
		{1},
	}
	fmt.Println(mysearch([]int{1,2,3},4))
	fmt.Println(searchMatrix(a,2))
}
