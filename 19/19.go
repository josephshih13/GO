package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr *ListNode
	l := 0
	for ptr = head; ptr != nil; ptr = ptr.Next {
		l++
	}
	tar := l - n
	if tar == 0 {
		return head.Next
	}
	ptr = head
	// var pre *ListNode
	for i := 0; i < tar-1; i++ {
		ptr = ptr.Next
	}
	ptr.Next = ptr.Next.Next
	// fmt.Println(l)
	return head
}

func main() {
	head1 := ListNode{Val: 1}
	//printNode(&head1)
	addNode(&head1, 2)
	addNode(&head1, 3)
	addNode(&head1, 4)
	addNode(&head1, 5)
	//addNode(&head1, 3)
	//printNode(&head1)
	eee := removeNthFromEnd(&head1, 5)
	// fmt.Println(eee)
	printNode(eee)
}
