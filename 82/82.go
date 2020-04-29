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
	if head == nil{
		return head
	}
	ma := make(map[int]bool)
	for pre,ptr:=1000000000,head;ptr!=nil;ptr=ptr.Next{
		if ptr.Val != pre{
			pre = ptr.Val
		} else{
			ma[ptr.Val] = true
		}
	}
	fmt.Println(ma)
	//ma[1]=true
	for head != nil{
		_,ex := ma[head.Val]
		if !ex{
			break
		}
		head = head.Next
	}
	if head == nil{
		return head
	}
	//printlist(head)
	for ptr:=head;ptr!=nil&&ptr.Next!=nil;ptr=ptr.Next{
		ptr2 := ptr
		for ptr2.Next != nil{
			_,ex := ma[ptr2.Next.Val]
			if ex {
				ptr2 = ptr2.Next
			} else{
				break
			}
		}
		ptr.Next = ptr2.Next
	}
	return head
}

func main() {
	a:=makelist([]int{1,1})
	a=deleteDuplicates(a)
	printlist(a)
	aa:=makelist([]int{1,1,1,2,3})
	aa=deleteDuplicates(aa)
	printlist(aa)
}
