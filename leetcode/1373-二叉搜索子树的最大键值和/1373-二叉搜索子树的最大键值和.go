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

// 1373-二叉搜索子树的最大键值和
func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := 0
	traverse(root, &maxVal)
	return maxVal
}

/*
定义⼀个 traverse 函数，traverse(root) 返回⼀个⼤⼩为 4 的 int 数组，我们暂且称它为 res，其 中：
res[0] 记录以 root 为根的⼆叉树是否是 BST，若为 1 则说明是 BST，若为 0 则说明不是 BST；
res[1] 记录以 root 为根的⼆叉树所有节点中的最⼩值；
res[2] 记录以 root 为根的⼆叉树所有节点中的最⼤值；
res[3] 记录以 root 为根的⼆叉树所有节点值之和。
*/
func traverse(root *TreeNode, maxVal *int) []int {
	if root == nil {
		return []int{1, math.MaxInt, math.MinInt, 0}
	}
	left := traverse(root.Left, maxVal)
	right := traverse(root.Right, maxVal)
	res := make([]int, 4)
	// 左子树，右子树是一个BST，加上当前结点也是BST
	if left[0] == 1 && right[0] == 1 && left[2] < root.Val && root.Val < right[1] {
		// 是BST树
		res[0] = 1
		res[1] = min(root.Val, left[1])
		res[2] = max(root.Val, right[2])
		res[3] = root.Val + right[3] + left[3]

		*maxVal = max(*maxVal, res[3])
	} else {
		// 不是BST，只需要标记，不需要进行计算
		res[0] = 0
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
