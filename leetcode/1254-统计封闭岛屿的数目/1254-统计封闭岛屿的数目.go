package main

// 1254-统计封闭岛屿的数目
func closedIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		// 把靠左边的岛屿淹掉
		dfs(grid, i, 0)
		// 把靠右边的岛屿淹掉
		dfs(grid, i, n-1)
	}
	for i := 0; i < n; i++ {
		// 把靠上边的岛屿淹掉
		dfs(grid, 0, i)
		// 把靠下边的岛屿淹掉
		dfs(grid, m-1, i)
	}
	// 遍历 grid，剩下的岛屿都是封闭岛屿
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海⽔
func dfs(grid [][]int, i, j int) {
	m := len(grid)
	n := len(grid[0])
	// 排除边缘 m,n 大于等于
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	// 已经是海⽔了
	if grid[i][j] == 1 {
		return
	}
	// 将 (i, j) 变成海⽔
	grid[i][j] = 1
	// 淹没上下左右的陆地
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}
