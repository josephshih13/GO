package main

import "fmt"

func removeDuplicates(nums []int) int {
	pre_num,pre_cnt := 10000000000,0
	idx:=0
	for i:=range nums{
		if nums[i]!=pre_num{
			pre_num = nums[i]
			pre_cnt = 0
		}
		if pre_cnt<2{
			nums[idx] = nums[i]
			idx++
		}
		pre_cnt++
	}
	return idx
}

func main() {
	a:=[]int{1,1,1,2,2,3}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
}
