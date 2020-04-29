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

func listlen(head *ListNode) int{
	ret :=0
	for head != nil{
		ret++
		head = head.Next
	}
	return ret
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := listlen(head)
	if l ==0{
		return head
	}
	k %= l
	k = l-k
	if k ==0 {
		return head
	}

	tmp := head
	for tmp.Next != nil{
		tmp = tmp.Next
	}
	tmp.Next = head

	new_head := head
	for i:=0;i<k;i++{
		new_head = new_head.Next
	}
	fmt.Println(new_head.Val)

	tmp = new_head
	for tmp.Next != new_head{
		tmp = tmp.Next
	}
	tmp.Next = nil

	return new_head
}

func main() {
	a := makelist([]int{1,2,3,4,5})
	printlist(a)
	a = rotateRight(a,2)
	printlist(a)
}
