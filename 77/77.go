package main

import (
	"fmt"
)

func dfs(now,tar int, arr,cond []int, used []bool,ans *[][]int){
	if now == tar{
		fmt.Println(arr)
		tmp:=make([]int,tar)
		copy(tmp,arr)
		*ans = append(*ans,tmp)
		return
	}
	for i:=0;i<len(cond);i++{
		if !used[i]&&(now==0||cond[i]>arr[now-1]){
			arr[now] = cond[i]
			used[i] = true
			dfs(now+1,tar,arr,cond,used,ans)
			used[i] = false
		}
	}
}

func combine(n int, k int) [][]int {
	cond:=[]int{}
	used := []bool{}
	arr:=[]int{}
	for i:=1;i<=n;i++{
		cond=append(cond,i)
		used=append(used,false)
	}
	for i:=0;i<k;i++{
		arr=append(arr,0)
	}
	ans:=[][]int{}

	dfs(0,k,arr,cond,used,&ans)
	return ans
}

func main() {
	fmt.Println(combine(4,2))
}
