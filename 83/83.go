package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(head *ListNode, num int) {
	if head.Next == nil {
		head.Next = &ListNode{Val: num}
		return
	}
	ptr := head.Next
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = &ListNode{Val: num}
}

func makelist(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head *ListNode
	tmp := ListNode{Val: nums[0]}
	head = &tmp
	for i := 1; i < len(nums); i++ {
		addNode(head, nums[i])
	}
	return head
}

func printlist(head *ListNode) {
	//fmt.Println(node.Val)
	//ptr := node.Next
	//if ptr != nil {
	//	printNode(ptr)
	//}
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println("")
}

func listlen(head *ListNode) int {
	ret := 0
	for head != nil {
		ret++
		head = head.Next
	}
	return ret
}

func deleteDuplicates(head *ListNode) *ListNode {
	ptr := head
	//pre := 10000000000
	for ptr!=nil{
		ptr2 := ptr.Next
		for ptr2 != nil && ptr2.Val == ptr.Val{
			ptr2 = ptr2.Next
		}
		ptr.Next = ptr2
		ptr = ptr2
	}
	return head
}

func main() {
	a:=makelist([]int{1,1,2})
	a = deleteDuplicates(a)
	printlist(a)
	aa:=makelist([]int{1,1,2,3,3})
	aa = deleteDuplicates(aa)
	printlist(aa)
}
