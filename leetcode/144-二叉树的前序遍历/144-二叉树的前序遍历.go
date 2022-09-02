package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144-二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
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
	(*res) = append((*res), root.Val)
	traversal(root.Left, res)
	traversal(root.Right, res)
}
