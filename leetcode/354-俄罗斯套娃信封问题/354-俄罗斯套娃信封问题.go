package main

import "sort"

// 算法超时，无法通过所有测试用例
// 354-俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	// 先对信封宽度排序，宽度相等对高度进行排序，之后对宽度进行求最长子序列
	sort.Slice(envelopes, func(i, j int) bool {
		// 按照宽度排序
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			// 宽度相等，高度降序
			if envelopes[i][1] > envelopes[j][1] {
				return true
			}
		}
		return false

	})
	height := make([]int, len(envelopes))
	for i := range envelopes {
		height[i] = envelopes[i][1]
	}
	return lengthOfLIS(height)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	dp := make([]int, len(nums)+1)
	// 每一个自己最少都是一个最长递增子序列，初始化为1
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	for i := range dp {
		res = max(res, dp[i])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
