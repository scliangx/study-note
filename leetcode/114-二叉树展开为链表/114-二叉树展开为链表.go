package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 114-二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 1、左右⼦树已经被拉平成⼀条链表
	left := root.Left
	right := root.Right

	// 2、将左⼦树作为右⼦树
	root.Right = left
	root.Left = nil

	// 3、将原先的右⼦树接到当前右⼦树的末端
	// 寻找原本左子树的最右结点，将原本二叉树的右子树连接到原本二叉树左子树的最右结点上
	p := root
	for root.Right != nil {
		p = p.Right
	}
	p.Right = right
}
