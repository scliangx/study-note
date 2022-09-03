package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104-二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 非递归的方式
func maxDepth2(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}
	return depth
}
