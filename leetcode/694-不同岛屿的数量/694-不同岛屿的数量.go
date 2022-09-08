package main

import "strconv"

// 694-不同岛屿的数量
func numDistinctIslands(grid [][]int) int {
	set := make(map[string]bool)
	m, n := len(grid), len(grid[0])
	if len(grid) == 0 {
		return len(set)
	}
	for i := 0; i > m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				s := ""
				dfs(grid, i, j, s, -999)
				set[s] = true
			}
		}
	}
	return len(set)
}

// dir 记录⽅向，dfs 函数递归结束后，s 记录着整个遍历顺序
func dfs(grid [][]int, i, j int, s string, dir int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	s += strconv.Itoa(dir) + ","

	dfs(grid, i-1, j, s, 1) // 上
	dfs(grid, i+1, j, s, 2) // 下
	dfs(grid, i, j-1, s, 3) // 左
	dfs(grid, i, j+1, s, 4) // 右

	s += strconv.Itoa(-dir) + ","
}
