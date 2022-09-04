package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 450-删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找到最右子树最小的结点或者最左子树的最大的结点
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 删除找到的结点，因为需要将找到的结点放在原本要删除的位置
		root.Right = deleteNode(root.Right, minNode.Val)
		// 更新移动上来的结点的左右子树
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
