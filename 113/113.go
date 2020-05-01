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

func hasPathSum(root *TreeNode, sum int) bool {
	if root ==nil{
		//if sum ==0{
		//	return true
		//}
		return false
	}
	var trav func(*TreeNode,int) bool
	trav = func(now *TreeNode,k int) bool{
		if now.Left == nil &&now.Right==nil&&k+now.Val==sum{
			return true
		}
		if now.Left!=nil && trav(now.Left,k+now.Val){
			return true
		}
		if now.Right!=nil&&trav(now.Right,k+now.Val){
			return true
		}
		return false
	}
	return trav(root,0)
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret :=[][]int{}
	if root==nil{
		return ret
	}
	aa := []int{}
	var trav func(*TreeNode,int,[]int)
	trav = func(now *TreeNode,k int,arr []int){
		if now.Left == nil &&now.Right==nil&&k+now.Val==sum{
			arr = append(arr,now.Val)
			tmp:=make([]int,len(arr))
			copy(tmp,arr)
			ret = append(ret,tmp)
			return
		}
		if now.Left!=nil{
			arr = append(arr,now.Val)
			trav(now.Left,k+now.Val,arr)
			arr = arr[:len(arr)-1]
		}
		if now.Right!=nil{
			arr = append(arr,now.Val)
			trav(now.Right,k+now.Val,arr)
			arr = arr[:len(arr)-1]
		}
	}
	trav(root,0,aa)
	return ret
}

func main() {

}
