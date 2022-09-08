package main

// 1905-统计子岛屿
// 思路： 淹没grid2 中是陆地，但是grid1中是海域的区块，然后再统计grid2中的陆地数量即可
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	if len(grid2) == 0 {
		return res
	}
	m, n := len(grid1), len(grid1[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果在grid1中是海水，在grid2中是岛屿，直接淹没
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(grid2, i, j)
			}
		}
	}
	// 现在 grid2 中剩下的岛屿都是⼦岛，计算岛屿数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
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
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i+1, j)

}
