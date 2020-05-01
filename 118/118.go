package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows ==0{
		return [][]int{}
	}
	ret := [][]int{}
	for i:=0;i<numRows;i++{
		ret = append(ret,make([]int,i+1))
	}
	ret[0][0]=1
	for i:=1;i<len(ret);i++{
		ret[i][0],ret[i][len(ret[i])-1]=1,1
		for j:=1;j<len(ret[i])-1;j++{
			ret[i][j] = ret[i-1][j-1]+ret[i-1][j]
		}
	}
	return ret
}

func main() {
	fmt.Println(generate(5))
}
