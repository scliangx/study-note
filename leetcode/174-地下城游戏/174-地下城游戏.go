package main

// 从公主的位置向骑士的位置走，每一个位置最低需要1点血量
// 174-地下城游戏
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				// 公主所在的位置，最低需要1点血，最高需要1-dungeon[i][j]
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				// 到了最后一行，那么最低需要1点血，最高需要前一列减去当前位置已有的血量
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				// 到了最后一列，那么最低需要1点血，最高需要下一列减去当前位置已有的血量
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				// 到达i行j列，最低需要前一列或者下一行中所需的最低血量减去当前位置已有的血量
				dp[i][j] = max(1, min(dp[i][j+1], dp[i+1][j])) - dungeon[i][j]
			}
		}
	}
	return dp[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
