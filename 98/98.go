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

func isValidBST(root *TreeNode) bool {
	if root==nil{
		return true
	}

	var trav func (*TreeNode,int,int) bool
	trav = func (now *TreeNode,min,max int) bool{
		//if now.Left!=nil && now.Left.Val >= now.Val{
		//	return false
		//}
		//if now.Right!=nil&&now.Right.Val <= now.Val{
		//	return false
		//}
		if now.Val >= max || now.Val<=min{
			return false
		}
		if now.Left != nil && !trav(now.Left,min,now.Val){
			return false
		}
		if now.Right !=nil&&!trav(now.Right,now.Val,max){
			return false
		}
		return true
	}
	return trav(root,-1<<60,1<<60)

}

func main() {

}
