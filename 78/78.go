package main

import "fmt"

func subsets(nums []int) [][]int {
	//used := make([]bool,len(nums))
	arr := []int{}
	ans := [][]int{}
	var dfs func (int)
	dfs = func(idx int) {
		//fmt.Println(idx)
		if idx == len(nums){
			tmp:=make([]int,len(arr))
			copy(tmp,arr)
			ans = append(ans,tmp)
			return
		}
		dfs(idx+1)
		arr = append(arr,nums[idx])
		dfs(idx+1)
		arr = arr[:len(arr)-1]
		//for i:=0;i<len(nums);i++{
		//	if !used[i]{
		//		used[i] = true
		//		arr = append(arr,nums[i])
		//		dfs(idx+1)
		//		arr = arr[:len(arr)-1]
		//		used[i] = false
		//	}
		//}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}
