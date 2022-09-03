package main

/**
* Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 105-从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0
	for i, val := range inorder {
		if val == preorder[0] {
			index = i
			break
		}
	}
	preMidd:= len(inorder[:index]) + 1
	root.Left = buildTree(preorder[1:preMidd], inorder[:index+1])
	root.Right = buildTree(preorder[preMidd:], inorder[index+1:])
	return root
}
