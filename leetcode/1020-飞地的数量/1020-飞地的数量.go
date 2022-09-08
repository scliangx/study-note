package main

// 1020-飞地的数量
func numEnclaves(grid [][]int) int {
	res := 0
	if len(grid) == 0 {
		return res
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		// 淹没掉所有的第0列的陆地
		dfs(grid, i, 0)
		// 淹没掉所有的第n-1列的陆地
		dfs(grid, i, n-1)
	}

	for i := 0; i < n; i++ {
		// 淹没掉所有的第0行的陆地
		dfs(grid, 0, i)
		// 淹没掉所有的第n-1列的陆地
		dfs(grid, m-1, i)
	}
	// 周边全部变成水之后，直接数⼀数剩下的陆地即可
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
}
