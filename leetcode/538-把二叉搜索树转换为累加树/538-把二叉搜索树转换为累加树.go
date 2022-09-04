package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 538-把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 更换递归遍历二叉搜索树左右子树的顺序，可以保证遍历结果为升序排序
	// 每次累加，遍历到当前累加的值一定是比自己大的值得和
	traverse(root.Right, sum)
	(*sum) += root.Val
	root.Val = (*sum)
	traverse(root.Left, sum)
}
