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

func mergelist(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ptr := l1
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = l2
	return l1
}

func partition(head *ListNode, x int) *ListNode {
	arr1,arr2:=[]int{},[]int{}
	for ptr:=head;ptr!=nil;ptr=ptr.Next{
		if ptr.Val<x{
			arr1 = append(arr1,ptr.Val)
		}else{
			arr2 = append(arr2,ptr.Val)
		}
	}
	l1:=makelist(arr1)
	l2:=makelist(arr2)
	l1=mergelist(l1,l2)
	return l1
}

func main() {

}
