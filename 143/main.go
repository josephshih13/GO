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

func revList(head *ListNode) *ListNode {
	pre, ptr := head, head
	pre = nil

	for ptr != nil {
		next := ptr.Next
		// nextnext := ptr.Next.Next
		ptr.Next = pre
		pre = ptr
		ptr = next
	}
	return pre
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil && slow.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	h2 := revList(slow.Next)
	slow.Next = nil
	// ret := head
	ptr := head
	h1 := head.Next
	for h2 != nil {
		ptr.Next = h2
		ptr = h2
		h2 = h2.Next
		if h1 != nil {
			ptr.Next = h1
			ptr = h1
			h1 = h1.Next
		}
	}
}
func main() {
	fmt.Println("123")
	s := "12345"
	fmt.Println(s[1:4] == "234")
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// [[1,2],[3,5],[6,7],[8,10],[12,16]]
// [4,8]
