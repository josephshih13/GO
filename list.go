package main

import "fmt"

type ListNode struct {
	Val int
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

func makelist(nums []int) *ListNode{
	if len(nums)==0 {
		return nil
	}
	var head *ListNode
	tmp := ListNode{Val: nums[0]}
	head = &tmp
	for i:=1;i<len(nums);i++{
		addNode(head,nums[i])
	}
	return  head
}

func printlist(head *ListNode) {
	//fmt.Println(node.Val)
	//ptr := node.Next
	//if ptr != nil {
	//	printNode(ptr)
	//}
	for head != nil{
		fmt.Print(head.Val,"->")
		head = head.Next
	}
	fmt.Println("")
}

func main() {
	l1:=makelist([]int{1,2,3,4})
	fmt.Println(l1)
	printlist(l1)
}
