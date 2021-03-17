/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func trav(node *Node, ma map[*Node]*Node) map[*Node]*Node{
    _, prs := ma[node]
    if prs{
        return ma
    }
    
    tmp := Node{Val:node.Val,Neighbors: []*Node{}}
    ma[node] = &tmp
    
    for _,v:=range(node.Neighbors){
        ma = trav(v,ma)
    }
    return ma
}

func copy_node(a,b *Node, ma map[*Node]*Node){
    for _,no := range(a.Neighbors){
        b.Neighbors = append(b.Neighbors,ma[no])
    }
}

func cloneGraph(node *Node) *Node {
    if node == nil{
        return nil
    }
    ma := map[*Node]*Node{}
    ma = trav(node,ma)
    for k,v := range(ma){
        copy_node(k,v,ma)
    }
    
    return ma[node]
    
}