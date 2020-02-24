package main

import (
	"fmt"
)

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

func printNode(node *ListNode) {
	fmt.Println(node.Val)
	ptr := node.Next
	if ptr != nil {
		printNode(ptr)
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head.Next
	head.Next, head.Next.Next = head.Next.Next, head
	ptr, pre := ans.Next.Next, ans.Next
	// fmt.Println("###")
	// printNode(ans)
	// fmt.Println("###")
	for ptr != nil && ptr.Next != nil {
		ptr.Next, ptr.Next.Next, pre.Next = ptr.Next.Next, ptr, ptr.Next
		pre = pre.Next.Next
		ptr = ptr.Next
		// fmt.Println("###")
		// printNode(ans)
		// // duration := time.Duration(5) * time.Second
		// // time.Sleep(duration)
		// fmt.Println(ptr, ptr.Next)
		// fmt.Println("###")
		// break
	}
	return ans
}

func main() {
	// head1 := ListNode{}
	//printNode(&head1)
	// addNode(&head1, 2)
	//addNode(&head1, 3)
	// addNode(&head1, 4)
	// printNode(&head1)

	head2 := ListNode{Val: 1}
	//printNode(&head1)
	addNode(&head2, 2)
	addNode(&head2, 3)
	addNode(&head2, 4)
	addNode(&head2, 5)
	printNode(&head2)

	//var test *ListNode
	//test = addTwoNumbers(&head1, &head2)
	//fmt.Println(test)
	printNode(swapPairs(&head2))
}
