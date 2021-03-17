/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isnull(x *TreeNode)bool  {
	if x ==nil {
		return true
	}
	return false
}

func trav(p,q *TreeNode) bool{
	if p.Val!=q.Val{
		return false
	}
	if isnull(p.Left) != isnull(q.Left){
		return false
	}
	if isnull(p.Right) != isnull(q.Right){
		return false
	}
	a,b:=true,true
	if !isnull(p.Left){
		a=trav(p.Left,q.Left)
	}
	if !isnull(p.Right){
		b=trav(p.Right,q.Right)
	}
	return a&&b
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if isnull(p)!=isnull(q){
		return false
	}
	if isnull(p)&&isnull(q){
		return true
	}
	return trav(p,q)
}

func trav2(s *TreeNode, t *TreeNode)bool{
    if isSameTree(s,t){
        return true
    }
    if !isnull(s.Left) && trav2(s.Left,t){
        return true
    }
    if !isnull(s.Right) && trav2(s.Right,t){
        return true
    }
    return false
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
    return trav2(s, t)
}