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

func printNode(node *ListNode) {
	fmt.Println(node.Val)
	ptr := node.Next
	if ptr != nil {
		printNode(ptr)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1, ptr2 := l1, l2
	var head *ListNode
	var ptr *ListNode
	addone := 0
	for ptr1 != nil || ptr2 != nil {
		var sum int
		if ptr1 == nil {
			tmp := ptr2.Val + addone
			sum = tmp % 10
			addone = tmp / 10
			ptr2 = ptr2.Next
		} else if ptr2 == nil {
			tmp := ptr1.Val + addone
			sum = tmp % 10
			addone = tmp / 10
			ptr1 = ptr1.Next
		} else {
			tmp := ptr1.Val + ptr2.Val + addone
			sum = tmp % 10
			addone = tmp / 10
			ptr1, ptr2 = ptr1.Next, ptr2.Next
		}
		if ptr == nil {
			ptr = &ListNode{Val: sum}
			head = ptr
		} else {
			ptr.Next = &ListNode{Val: sum}
			ptr = ptr.Next
		}
	}
	if addone != 0 {
		ptr.Next = &ListNode{Val: addone}
	}
	return head
}

func main() {
	head1 := ListNode{Val: 9}
	//printNode(&head1)
	addNode(&head1, 9)
	//addNode(&head1, 3)
	printNode(&head1)

	head2 := ListNode{Val: 1}
	//printNode(&head1)
	//addNode(&head2, 6)
	//addNode(&head2, 4)
	printNode(&head2)

	//var test *ListNode
	//test = addTwoNumbers(&head1, &head2)
	//fmt.Println(test)
	printNode(addTwoNumbers(&head1, &head2))

}
