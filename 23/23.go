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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ptr1, ptr2, ans, now *ListNode
	ptr1, ptr2 = l1, l2
	for ptr1 != nil || ptr2 != nil {
		// printNode(ptr1)
		// printNode(ptr2)
		var small int
		if ptr2 == nil || (ptr1 != nil && ptr1.Val < ptr2.Val) {
			small = ptr1.Val
			ptr1 = ptr1.Next
		} else {
			small = ptr2.Val
			ptr2 = ptr2.Next
		}
		// fmt.Println("??", small)
		if ans == nil {
			tmp := ListNode{}
			tmp.Val = small
			ans, now = &tmp, &tmp
		} else {
			tmp := ListNode{}
			tmp.Val = small
			now.Next = &tmp
			now = now.Next
			now.Val = small
		}
		// fmt.Println("###")
		// printNode(ans)
		// printNode(now)
		// fmt.Println(ptr1, ptr2)
		// fmt.Println("###")
	}
	return ans
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	k := lists[0]
	for i := 1; i < len(lists); i++ {
		k = mergeTwoLists(lists[i], k)
	}
	return k
}

func main() {
	head1 := ListNode{Val: 1}
	//printNode(&head1)
	addNode(&head1, 4)
	//addNode(&head1, 3)
	addNode(&head1, 5)
	// printNode(&head1)

	head2 := ListNode{Val: 1}
	//printNode(&head1)
	addNode(&head2, 3)
	addNode(&head2, 4)
	// printNode(&head2)

	head3 := ListNode{Val: 2}
	addNode(&head3, 6)

	lists := []*ListNode{&head1, &head2, &head3}

	//var test *ListNode
	//test = addTwoNumbers(&head1, &head2)
	//fmt.Println(test)
	// printNode(mergeTwoLists(nil, &head2))
	printNode(mergeKLists(lists))

}
