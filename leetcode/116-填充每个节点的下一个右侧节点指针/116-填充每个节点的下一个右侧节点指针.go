package main

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116-填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 递归思路
func connect2(root *Node) *Node {
    if root==nil{
        return nil
    }
    connectTwoNode(root.Left,root.Right)
    return root 
}

func connectTwoNode(node1 *Node,node2 *Node){
    if node1==nil||node2==nil{
        return
    }
    node1.Next=node2
    connectTwoNode(node1.Left,node1.Right)
    connectTwoNode(node2.Left,node2.Right)
    connectTwoNode(node1.Right,node2.Left)
}

