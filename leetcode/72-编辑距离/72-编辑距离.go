package main

// 72-编辑距离
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果word1和word2当前位置相等，则直接下一个
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 否则的话在增删改其中选择操作最少的
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(val ...int) int {
	minVal := val[0]
	for i := 1; i < len(val); i++ {
		if val[i] < minVal {
			minVal = val[i]
		}
	}
	return minVal
}
