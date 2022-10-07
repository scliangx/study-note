package main

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 103-二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
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
	for i := 1; i < len(res); i += 2 {
		res[i] = reverse(res[i])
	}
	return res
}

func reverse(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}
