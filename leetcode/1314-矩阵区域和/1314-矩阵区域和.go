package main

// 1314-矩阵区域和
func matrixBlockSum(mat [][]int, k int) [][]int {
	matrix := Constructor(mat)
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x1 := max(i-k, 0)
			y1 := max(j-k, 0)
			x2 := min(i+k, m-1)
			y2 := min(j+k, n-1)
			res[i][j] = matrix.SumRegion(x1, y1, x2, y2)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}
	res := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 加上上、左、左上;减去左上重复的
			res[i][j] = res[i-1][j] + res[i][j-1] + matrix[i-1][j-1] - res[i-1][j-1]
		}
	}
	return NumMatrix{
		preSum: res,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}
