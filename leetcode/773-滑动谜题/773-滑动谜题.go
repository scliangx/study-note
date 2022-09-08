package main

import "strconv"

// 773-滑动谜题
// ----------------------------***************---------------------------
// 如果⼆维数组中的某个元素 e 在⼀维数组中的索引为 i，
// 那么 e 的左右相邻元素在⼀维 数组中的索引就是 i - 1 和 i + 1，
// ⽽ e 的上下相邻元素在⼀维数组中的索引就是 i - n 和 i + n，
//其中 n 为⼆维数组的列数。 这样，对于 m x n 的⼆维数组
func slidingPuzzle(board [][]int) int {
	// 维持一个二维到一维数组下标的映射，分别是他的上下左右的位置
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	m, n := 2, 3
	target := "123450"
	sb := ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sb += strconv.Itoa(board[i][j])
		}
	}
	var queue []string
	step := 0
	visited := make(map[string]bool)
	visited[sb] = true
	queue = append(queue, sb)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			// 找到可以动的位置
			idx := 0
			for cur[idx] != '0' {
				idx++
			}
			for _, adj := range neighbor[idx] {
				newBoard := swap(cur, adj, idx)
				if !visited[newBoard] {
					queue = append(queue, newBoard)
					visited[newBoard] = true
				}
			}
		}
		step++
	}
	return -1
}

func swap(s string, i, j int) string {
	c := []int32{}
	for _, v := range s {
		c = append(c, v)
	}
	c[i], c[j] = c[j], c[i]
	return string(c)
}
