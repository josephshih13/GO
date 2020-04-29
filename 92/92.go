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

func reverseList(head *ListNode) *ListNode {
	var ret, pre *ListNode
	pre = nil
	for ptr := head; ptr != nil; {
		if ptr.Next == nil {
			ret = ptr
		}
		tmp := ptr.Next
		ptr.Next = pre
		pre = ptr
		ptr = tmp
	}
	return ret
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	ptr:=head
	idx:=1
	for idx != n{
		idx++
		ptr =ptr.Next
	}
	tmp:=ptr.Next
	ptr.Next=nil
	ptr = head
	idx=1
	for idx!=m{
		idx++
		ptr=ptr.Next
	}
	//printlist(head)
	h:=reverseList(ptr)
	//printlist(h)
	if m==1{
		head = h
	}else{
		pp:=head
		ii:=1
		for ii!=m-1{
			pp=pp.Next
			ii++
		}
		pp.Next = h
	}
	//printlist(head)
	pp:=head
	for pp.Next!=nil{
		pp=pp.Next
	}
	pp.Next = tmp
	return head
}

func main() {
	a:=makelist([]int{1,2,3,4,5})
	a=reverseBetween(a,2,4)
	printlist(a)
}
