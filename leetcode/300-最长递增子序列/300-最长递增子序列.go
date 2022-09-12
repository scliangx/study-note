package main

// 300-最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	// 定义：dp[i] 表示以 nums[i] 这个数结尾的最⻓递增⼦序列的⻓度
	dp := make([]int, len(nums)+1)
	// base case：dp 数组全都初始化为 1
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 寻找 nums[0..j-1] 中⽐ nums[i] ⼩的元素
			if nums[j] < nums[i] {
				// 把 nums[i] 接在后⾯，即可形成⻓度为 dp[j] + 1，
				// 且以 nums[i] 为结尾的递增⼦序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

	}
	for i := range dp {
		if res < dp[i] {
			res = max(dp[i], res)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
