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

func levelOrder(root *TreeNode) [][]int {
	h := maxDepth(root)
	ans := [][]int{}
	if root==nil{
		return ans
	}
	for i:=0;i<h;i++{
		ans = append(ans,[]int{})
	}
	var trav func(*TreeNode,int)
	trav = func(now *TreeNode, deep int){
		if now.Left!=nil{
			trav(now.Left,deep+1)
		}
		ans[deep] = append(ans[deep],now.Val)
		if now.Right!=nil{
			trav(now.Right,deep+1)
		}
	}
	trav(root,0)
	return ans
}

func rev(a [][]int) [][]int{
	r := [][]int{}
	for i:=0;i<len(a);i++{
		r=append(r,[]int{})
	}
	for i:=0;i<len(a);i++{
		r[i] = a[len(a)-1-i]
	}
	return r
}

func levelOrderBottom(root *TreeNode) [][]int {
	tmp:=levelOrder(root)
	return rev(tmp)
}

func main() {

}
