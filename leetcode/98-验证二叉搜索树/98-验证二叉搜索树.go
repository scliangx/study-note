package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98-验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBSTHelper(root, nil, nil)
}

func isBSTHelper(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	// 右子树的值不能小于父结点的值
	if min != nil && root.Val <= min.Val {
		return false
	}
	// 左子树不能大于父结点的值
	if max != nil && root.Val >= max.Val {
		return false
	}

	return isBSTHelper(root.Left, min, root) && isBSTHelper(root.Right, root, max)
}
