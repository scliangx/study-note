package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 501-二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	// 基于递归的有效变量维护
	maxCnt := 0 // 最大频次
	count := 0  // 当前元素的频次
	pre := &TreeNode{}
	res := []int{} // 众数
	// 中序遍历
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// 与前驱节点同，cnt++
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		// 该数的的频次，与全局的最大频次同，加入结果集
		if count == maxCnt {
			res = append(res, node.Val)
		}
		// 遇到更大频次，需要更新频次，清空res，加入res
		if count > maxCnt {
			res = []int{}
			maxCnt = count
			res = append(res, node.Val)
		}
		// 更新前驱节点
		pre = node
		traverse(node.Right)

	}
	traverse(root)
	return res
}
