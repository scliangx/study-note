package main

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 95-不同的二叉搜索树II
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	ans := []*TreeNode{}
	for i := start; i <= end; i++ {
		// 比当前值小的构造左子树
		leftTree := generate(start, i-1)
		// 比当前值大的构造右子树
		rightTree := generate(i+1, end)
		// 分辨对左右子树构建
		for _, left := range leftTree {
			for _, right := range rightTree {
				ans = append(ans, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return ans
}
