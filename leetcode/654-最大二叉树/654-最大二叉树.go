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

// 654-最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	// 寻找最大值，记录索引
	index := -1
	maxVal := math.MinInt
	for i := l; i <= r; i++ {
		if nums[i] > maxVal {
			index = i
			maxVal = nums[i]
		}
	}
	node := &TreeNode{Val: maxVal}
	node.Left = buildTree(nums, l, index-1)
	node.Right = buildTree(nums, index+1, r)
	return node
}
