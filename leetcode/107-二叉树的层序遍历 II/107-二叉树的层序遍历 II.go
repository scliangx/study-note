package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 107-二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}

	l, r := 0, len(res)
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
