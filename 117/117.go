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

func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	var trav func(*TreeNode)
	trav = func(now *TreeNode) {
		ret = append(ret, now.Val)
		if now.Left != nil {
			trav(now.Left)
		}

		if now.Right != nil {
			trav(now.Right)
		}
	}
	trav(root)
	return ret
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

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	max := -1
	var trav func(*Node, int)
	trav = func(now *Node, deep int) {
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

func levelOrder2(root *Node) [][]*Node {
	h := maxDepth2(root)
	ans := [][]*Node{}
	if root==nil{
		return ans
	}
	for i:=0;i<h;i++{
		ans = append(ans,[]*Node{})
	}
	var trav func(*Node,int)
	trav = func(now *Node, deep int){
		if now.Left!=nil{
			trav(now.Left,deep+1)
		}
		ans[deep] = append(ans[deep],now)
		if now.Right!=nil{
			trav(now.Right,deep+1)
		}
	}
	trav(root,0)
	return ans
}

func connect(root *Node) *Node {
	h := maxDepth2(root)
	ref := levelOrder2(root)
	fmt.Println(ref)
	for i:=0;i<h;i++{
		for j:=0;j<len(ref[i])-1;j++{
			ref[i][j].Next = ref[i][j+1]
		}
	}
	return root
}


func main() {

}
