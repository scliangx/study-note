package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102-二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		tmp := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
