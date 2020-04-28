package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	ma := make(map[int]int)
	for _,v:=range nums{
		_,ex := ma[v]
		if !ex{
			ma[v] = 1
		} else{
			ma[v]++
		}
	}
	//fmt.Println(ma)
	ks := []int{}
	for k:=range ma{
		ks = append(ks,k)
	}
	//fmt.Println(ks)
	ans := [][]int{}
	arr := []int{}
	var dfs func(int)
	dfs = func (idx int){
		if idx == len(ks){
			//fmt.Println(arr)
			tmp := make([]int,len(arr))
			copy(tmp,arr)
			ans = append(ans,tmp)
			return
		}
		for i:=0;i<=ma[ks[idx]];i++{
			for j:=0;j<i;j++{
				arr = append(arr,ks[idx])
			}
			dfs(idx+1)
			arr = arr[:len(arr)-i]
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(subsetsWithDup([]int{1,2,2}))
}
