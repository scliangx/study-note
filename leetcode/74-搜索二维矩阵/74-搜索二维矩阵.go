package main

// 74-搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		if getVal(matrix, mid) == target {
			return true
		} else if getVal(matrix, mid) > target {
			right = mid - 1
		} else if getVal(matrix, mid) < target {
			left = mid + 1
		}
	}
	return false
}

func getVal(matrix [][]int, index int) int {
	n := len(matrix[0])
	// 根据一维索引算出二维索引，找出数据
	i := index / n
	j := index % n
	return matrix[i][j]
}
