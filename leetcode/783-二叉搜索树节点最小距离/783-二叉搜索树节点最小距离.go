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

// 783-二叉搜索树节点最小距离
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minVal := math.MaxInt64
	pre := -1
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if pre != -1 && root.Val-pre < minVal {
			minVal = root.Val - pre
		}
		pre = root.Val
		traverse(root.Right)
	}
	traverse(root)
	return minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
