package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 965-单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val
	isUnival := true
	traverse(root, val, &isUnival)
	return isUnival
}

func traverse(root *TreeNode, val int, isUnival *bool) {
	if root == nil || !(*isUnival) {
		return
	}
	if root.Val != val {
		*isUnival = false
		return
	}
	traverse(root.Left, val, isUnival)
	traverse(root.Right, val, isUnival)
}
