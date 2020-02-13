package main

import "fmt"

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	cnt := 0
	for idx, num := range nums {
		if target%2 == 0 && num == target/2 {
			cnt++
		} else {
			set[num] = idx
		}
	}
	if cnt >= 2 {
		ans1, ans2 := -1, -1
		//fmt.Println(ans1 != -1)
		for idx, num := range nums {
			if num == target/2 {
				if ans1 == -1 {
					ans1 = idx
				} else {
					ans2 = idx
					//fmt.Println("??")
					break
				}
			}
			//fmt.Println(ans1, " ", ans2)
		}
		return []int{ans1, ans2}
	}
	//var ans []int
	ans := make([]int, 2)
	for idx, num := range nums {
		diff := target - num
		//fmt.Println(diff)
		tmp, appear := set[diff]
		if appear == true {
			ans[0] = idx
			ans[1] = tmp
			break
		}
	}
	return ans
}

func main() {
	arr := []int{3, 3, 9}
	fmt.Println(arr)
	tar := 6
	fmt.Println(twoSum(arr, tar))
}
