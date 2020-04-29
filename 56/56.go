package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	a,b int
}

type PA []Pair

func (p_arr PA) Len() int{
	return len(p_arr)
}

func (p_arr PA) Less(i,j int) bool{
	return p_arr[i].a < p_arr[j].a || p_arr[i].a==p_arr[j].a && p_arr[i].b < p_arr[j].b
}

func (p_arr PA) Swap(i,j int) {
	p_arr[i],p_arr[j] = p_arr[j],p_arr[i]
}
func maxint(a,b int) int {
	if a>b{
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals)==0{
		return [][]int{}
	}
	p_arr := PA{}
	for i:=range intervals{
		tmp := Pair{}
		tmp.a,tmp.b = intervals[i][0],intervals[i][1]
		p_arr = append(p_arr,tmp)
	}
	fmt.Println(p_arr)
	sort.Sort(p_arr)
	fmt.Println(p_arr)
	ret := [][]int{}
	aa,bb:=p_arr[0].a,p_arr[0].b
	for i:=1;i<p_arr.Len();i++{
		if p_arr[i].a>bb{
			tmp:=[]int{aa,bb}
			ret = append(ret,tmp)
			aa,bb = p_arr[i].a,p_arr[i].b
		} else{
			bb = maxint(bb,p_arr[i].b)
		}
	}
	tmp:=[]int{aa,bb}
	ret = append(ret,tmp)
	return ret
}

func main() {
	a:=[][]int{
		{1,4},
		{4,5},
	}
	b := merge(a)
	fmt.Println(b)
	aa:=[][]int{
		{1,3},
		{2,6},
		{8,10},
		{15,18},
	}
	bb:=merge(aa)
	fmt.Println(bb)
}
