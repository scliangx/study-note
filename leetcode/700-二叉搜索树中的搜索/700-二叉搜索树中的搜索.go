package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 700-二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return root
	}
}
