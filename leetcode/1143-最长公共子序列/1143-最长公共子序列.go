package main

// 1143-最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 创建 m+1 行 n+1 列的二维数组dp ，
	//其中 dp[i][j] 表示 text1 [0:i] 和 text2[0:j] 的最长公共子序列的长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 因此动态规划的边界情况是：当 i=0 或 j=0 时，dp[i][j]=0
	for i, val1 := range text1 {
		for j, val2 := range text2 {
			if val1 == val2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
