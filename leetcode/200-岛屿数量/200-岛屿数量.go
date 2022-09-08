package main

// 200-岛屿数量
func numIslands(grid [][]byte) int {
	res := 0
	if len(grid) == 0 {
		return res
	}
	m, n := len(grid), len(grid[0])
	var dfs func([][]byte, int, int)
	dfs = func(g [][]byte, i, j int) {
		m1, n1 := len(grid), len(grid[0])
		// 超出索引边界
		if i < 0 || j < 0 || i >= m1 || j >= n1 {
			return
		}
		// 已经是海⽔了
		if g[i][j] == '0' {
			return
		}
		// 从 (i, j) 开始，将与之相邻的陆地都变成海⽔
		// 淹没上下左右的陆地
		g[i][j] = '0'
		dfs(g, i-1, j)
		dfs(g, i+1, j)
		dfs(g, i, j-1)
		dfs(g, i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 每发现⼀个岛屿，岛屿数量加⼀
				res++
				// 然后使⽤ DFS 将岛屿淹了
				dfs(grid, i, j)
			}
		}
	}
	return res
}
