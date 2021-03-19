/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil{
        return nil
    }
    ma := map[*ListNode]bool{}
    ptr := headA
    for ptr != nil{
        ma[ptr] = true
        ptr = ptr.Next
    }
    ptr = headB
    for ptr != nil{
        _, prs := ma[ptr]
        if prs{
            return ptr
        }
        ptr = ptr.Next
    }
    return nil
}