package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230-二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	if root == nil {
		return res
	}
	traverse(root, &k, &res)
	return res
}

// 二叉搜索树的中序遍历是有序的
func traverse(root *TreeNode, k, res *int) {
	if root == nil {
		return
	}
	traverse(root.Left, k, res)
	(*k)--
	if (*k) == 0 {
		(*res) = root.Val
	}
	traverse(root.Right, k, res)
}
