package main

// 416-分割等和子集
func canPartition(nums []int) bool {
	sumVal := 0
	for _, val := range nums {
		sumVal += val
	}
	if sumVal%2 != 0 {
		return false
	}
	sum := sumVal / 2
	// 创建二维数组 dp，包含 n 行sum+1 列，其中dp[i][j] 表示从数组的 [0,i]下标范围内选取若干个正整数（可以是 0 个），
	// 是否存在一种选取方案使得被选取的正整数的和等于 j
	dp := make([][]bool, len(nums)+1)
	// 如果不选取任何正整数，则被选取的正整数等于 0,因此对于所有 0≤i<n，都有 dp[i][0]=true。
	// 当 i==0时，只有一个正整数 nums[0] 可以被选取，因此 dp[0][nums[0]]=true。
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= len(nums); i++ {
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			// 如果j - nums[i-1] < 0 : 则表示无法在选了，不然会超过sum
			if j-nums[i-1] < 0 {
				// 从nums[0...i-1]选择和为j是否可行
				dp[i][j] = dp[i-1][j]
			} else {
				// // 从nums[0...i-1]选择和为j是否可行 或者从nums[0...i-1] 选择和为j-nums[i-1]的是否可行
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}
