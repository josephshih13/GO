package main

import "fmt"

func sortColors(nums []int)  {
	arr := [3]int{}
	for _,v:=range nums{
		arr[v]++
	}
	fmt.Println(arr)
	for i,j:=0,0;i<len(nums);i++{
		for arr[j]==0{
			j++
		}
		nums[i]=j
		arr[j]--
	}
}

func main() {
	a:=[]int{2,0,2,1,1,0}
	sortColors(a)
	fmt.Println(a)
}
