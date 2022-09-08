package main

// 695-岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	maxRes := 0
	if len(grid) == 0 {
		return maxRes
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxRes = max(maxRes, dfs(grid, i, j))
			}
		}
	}
	return maxRes
}

// 淹没与 (i, j) 相邻的陆地，并返回淹没的陆地⾯积
func dfs(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs(grid, i, j-1) + dfs(grid, i, j+1) + dfs(grid, i-1, j) + dfs(grid, i+1, j) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
