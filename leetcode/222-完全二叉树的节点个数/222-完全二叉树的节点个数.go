package main

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 222-完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	leftDepth, rightDepth := 0, 0
	for l != nil {
		l = l.Left
		leftDepth++
	}
	for r != nil {
		r = r.Right
		rightDepth++
	}
	if leftDepth == rightDepth {
		return int(math.Pow(2, float64(leftDepth)) - 1)
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
