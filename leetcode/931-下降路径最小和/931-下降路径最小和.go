package main

import "math"

// 931-下降路径最小和
// 暴力递归解法
func minFallingPathSum1(matrix [][]int) int {
	res := math.MaxInt
	size := len(matrix)

	for i := 0; i < size; i++ {
		res = min(res, dp(matrix, size-1, i))
	}
	return res
}

func dp(matrix [][]int, i, j int) int {
	// 边界值不合法返回一个特殊值
	if i < 0 || j < 0 || i > len(matrix) || j > len(matrix[0]) {
		return 99999
	}
	// 到达第一行的时候开始返回
	if i == 0 {
		return matrix[i][j]
	}
	// 从最后一行往上搜索，搜索到第一个行结束，找到搜索的最小值
	return matrix[i][j] + min(min(dp(matrix, i-1, j), dp(matrix, i-1, j-1)), dp(matrix, i-1, j-1))

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 931-下降路径最小和
// 动态规划，备忘录 
func minFallingPathSum(matrix [][]int) int {
	res := math.MaxInt
	// 从第一行开始到最后一行结束
	dpTable := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dpTable[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); i++ {
			if i == 0 {
				// 第一行的最小值就是原始的值，直接赋值
				dpTable[i][j] = matrix[i][j]
			} else if j == 0 {
				// 如果是第一列，只有上方和右上
				dpTable[i][j] = min(dpTable[i-1][j], dpTable[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix)-1 {
				// 如果是最后一列，只有上方和左上
				dpTable[i][j] = min(dpTable[i-1][j-1], dpTable[i-1][j]) + matrix[i][j]
			} else {
				dpTable[i][j] = min(min(dpTable[i-1][j-1], dpTable[i-1][j]), dpTable[i-1][j+1]) + matrix[i][j]
			}
			if i == len(matrix)-1 && dpTable[i][j] < res {
				res = dpTable[i][j]
			}
		}

	}
	return res
}
