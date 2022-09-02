package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94-二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, res)
	(*res) = append((*res), root.Val)
	traversal(root.Right, res)
}
