package main

/**
 * Definition for a binary tree node.

 */
type Node struct {
	Val      int
	Children []*Node
}

// 559-N 叉树的最大深度
// 广度优先
func maxDepth(root *Node) int {
	depth := 0
	if root == nil {
		return depth
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, v := range node.Children {
				queue = append(queue, v)
			}
		}
		depth++
	}
	return depth
}

// 深度优先
func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		childDepth := maxDepth(child)
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}
