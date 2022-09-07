package main

// 51-N 皇后
func solveNQueens(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
	}
	// 初始化一个棋盘
	board := [][]string{}
	for i := 0; i < n; i++ {
		row := []string{}
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	backtrackQ(board, 0, &res)
	return res
}

func backtrackQ(board [][]string, row int, res *[][]string) {
	// 如果到达最后一行,return
	if row == len(board) {
		// 当前是一种结果，但是返回的是一个字符串数组所以需要进行转换
		*res = append(*res, covert(board))
		return
	}
	// 对每一行进行backtrack
	for col := 0; col < len(board[row]); col++ {
		// 当前位置合法进行放置
		if isVaild(board, row, col) {
			board[row][col] = "Q"
			backtrackQ(board, row+1, res)
			board[row][col] = "."
		}
	}
}

// 因为皇后是⼀⾏⼀⾏从上往下放的，所以左下⽅，右下⽅和正下⽅不⽤检查（还没放皇后）；
// 因为⼀ ⾏只会放⼀个皇后，所以每⾏不⽤检查。也就是最后只⽤检查上⾯，左上，右上三个⽅向
func isVaild(board [][]string, row int, col int) bool {
	// 检查列是否有皇后互相冲突
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == "Q" {
			return false
		}
	}
	//  检查右上⽅是否有皇后互相冲突
	m, n := row-1, col+1
	for m >= 0 && n < len(board) {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n++
	}
	//  检查左上⽅是否有皇后互相冲突
	m, n = row-1, col-1
	for m >= 0 && n >= 0 {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n--
	}
	return true
}

// 将二维数组转换成结果形式
// [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
func covert(board [][]string) []string {
	res := []string{}
	for i := 0; i < len(board); i++ {
		row := ""
		for j := 0; j < len(board[i]); j++ {
			row += board[i][j]
		}
		res = append(res, row)
	}
	return res
}
