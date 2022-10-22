package main

// 96-不同的二叉搜索树
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return count(1, n, &dp)
}

func count(low, high int, dp *[][]int) int {
	if low > high {
		return 1
	}
	if (*dp)[low][high] != 0 {
		return (*dp)[low][high]
	}
	res := 0
	for i := low; i <= high; i++ {
		// 以当前结点为根，左右子树分别有多少种
		left := count(low, i-1, dp)
		right := count(i+1, high, dp)
		res += (left * right)
	}
	(*dp)[low][high] = res
	return res
}
