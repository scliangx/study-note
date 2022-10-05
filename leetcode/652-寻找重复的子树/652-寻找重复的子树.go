package main

import (
	"strconv"
)

// 652-寻找重复的子树

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 652-寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	// 将树的路径序列化存储
	memo := make(map[string]int)
	res := []*TreeNode{}

	traverse(root, memo, &res)
	return res
}

func traverse(root *TreeNode, memo map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left, memo, res)
	right := traverse(root.Right, memo, res)

	subTree := left + "," + right + "," + strconv.Itoa(root.Val)
	// 记录子树出现的次数
	memo[subTree]++
	// 有重复的添加到答案中
	if memo[subTree] == 2 {
		*res = append(*res, root)
	}

	return subTree
}
