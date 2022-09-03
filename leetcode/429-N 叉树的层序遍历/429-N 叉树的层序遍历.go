package main

/**
 * Definition for a binary tree node.

 */
type Node struct {
	Val      int
	Children []*Node
}

// 429-N 叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			// 遍历多叉树的所有子结点
			for _, childNode := range node.Children {
				if childNode != nil {
					queue = append(queue, childNode)
				}
			}
		}
		res = append(res, tmp)
	}
	return res
}
