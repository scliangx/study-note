package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 543-二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	if root == nil {
		return maxDiameter
	}
	// 获取子树的最大深度
	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := dfs(tn.Left)
		right := dfs(tn.Right)
		maxDiameter = max(maxDiameter, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return maxDiameter
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
