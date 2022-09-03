package main

import "math"

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 515-在每个树行中找最大值
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		maxVal := math.MinInt
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxVal {
				maxVal = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, maxVal)
	}
	return res
}
