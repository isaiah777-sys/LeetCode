package leetcode

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]*Node)
	var clone func(*Node) *Node
	clone = func(n *Node) *Node {
		if v, ok := visited[n.Val]; ok {
			return v
		}
		newNode := &Node{Val: n.Val}
		visited[n.Val] = newNode
		for _, neighbor := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clone(neighbor))
		}
		return newNode
	}
	return clone(node)
}

type Node struct {
	Val       int
	Neighbors []*Node
}
