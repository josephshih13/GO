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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := -1
	var trav func(*TreeNode, int)
	trav = func(now *TreeNode, deep int) {
		if deep > max {
			max = deep
		}
		if now.Right != nil {
			trav(now.Right, deep+1)
		}
		if now.Left != nil {
			trav(now.Left, deep+1)
		}
	}
	trav(root, 1)
	return max
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := 1<<60
	var trav func(*TreeNode, int)
	trav = func(now *TreeNode, deep int) {
		if deep < min && now.Right == nil && nil == now.Left{
			min = deep
		}
		if now.Right != nil {
			trav(now.Right, deep+1)
		}
		if now.Left != nil {
			trav(now.Left, deep+1)
		}
	}
	trav(root, 1)
	return min
}

func main() {

}
