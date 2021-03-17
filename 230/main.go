/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func trav(root *TreeNode, k int, ans []int) []int{
    if root == nil{
        return ans
    }
    ans = trav(root.Left, k , ans)
    if len(ans) == k {
        return ans
    }
    ans = append(ans, root.Val)
    if len(ans) == k {
        return ans
    }
    ans = trav(root.Right,k,ans)
    return ans
}

func kthSmallest(root *TreeNode, k int) int {
    sl := trav(root, k , []int{})
    return sl[len(sl)-1]
}