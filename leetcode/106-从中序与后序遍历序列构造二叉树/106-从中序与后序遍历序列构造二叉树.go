package main

/**
* Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 106-从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	index := 0
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[:index+1], postorder[:index+1])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
