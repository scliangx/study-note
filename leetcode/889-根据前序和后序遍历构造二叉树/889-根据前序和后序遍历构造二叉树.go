package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 889-根据前序和后序遍历构造二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}
	// 根据前序遍历构建根结点
	root := &TreeNode{Val: postorder[0]}
	index := 0
	for i, val := range postorder {
		if val == preorder[1] {
			index = i
			break
		}
	}
	// 后续遍历第一个元素是左子树的导数第二个元素
	root.Left = constructFromPrePost(preorder[1:index+2], postorder[:index+1])
	root.Right = constructFromPrePost(preorder[index+2:], postorder[index+1:len(postorder)-1])
	return root
}
