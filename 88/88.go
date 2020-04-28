package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	tmp := make([]int,len(nums1))
	for i1,i2,i3:=0,0,0;i1<m||i2<n;{
		if i1 >=m{
			tmp[i3] = nums2[i2]
			i3++
			i2++
		}else if i2>=n{
			tmp[i3] = nums1[i1]
			i1++
			i3++
		}else if nums1[i1] <= nums2[i2]{
			tmp[i3] = nums1[i1]
			i1++
			i3++
		} else {
			tmp[i3] = nums2[i2]
			i3++
			i2++
		}
	}
	copy(nums1,tmp)

}

func main() {
	n1:=[]int{1,2,3,0,0,0}
	n2:=[]int{2,5,6}
	merge(n1,3,n2,3)
	fmt.Println(n1)
}
