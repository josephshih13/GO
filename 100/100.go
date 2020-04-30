package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

//func (x *TreeNode) isnull()bool {
//	if x ==nil {
//		return true
//	}
//	return false
//}

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

func main() {

}
