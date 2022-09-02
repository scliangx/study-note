package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 145-二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
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
	traversal(root.Right, res)
    (*res) = append((*res), root.Val)
}
