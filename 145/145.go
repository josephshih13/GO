package main

import "fmt"

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

func postorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil{
		return ret
	}
	var trav func (*TreeNode)
	trav = func (now *TreeNode){

		if now.Left!=nil{
			trav(now.Left)
		}

		if now.Right!=nil{
			trav(now.Right)
		}
		ret = append(ret,now.Val)
	}
	trav(root)
	return ret
}

func main() {
	a:=Ints2TreeNode([]int{1,NULL,2,3})
	fmt.Println(postorderTraversal(a))
}
