package main

import "math"

// 53-最大子数组和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := math.MinInt
	// 以i结尾的子数组的和是dp[i]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 每一个取当前结尾的子数组的和喝前一个值得子数组的和的最大值即可
		// 如果遇见负数
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	// dp中的最大值就是最大的子数组和
	for i := range dp {
		res = max(dp[i], res)
	}
	return res

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
