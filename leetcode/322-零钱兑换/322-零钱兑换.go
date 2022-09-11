package main

import (
	"math"
)

// 322-零钱兑换
// 暴力递归
func coinChange(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt
	for _, v := range coins {
		subRes := coinChange(coins, amount-v)
		// 等于-1表示当前凑不出来
		if subRes == -1 {
			continue
		}
		// 在⼦问题中选择最优解，然后加⼀,每一次选择了一颗硬币
		res = min(res, subRes+1)
	}
	if res == math.MaxInt {
		return -1
	}
	return coinChange(coins, amount)
}

// 定义：要凑出⾦额 n，⾄少要 dp(coins, n) 个硬币
// 备忘录
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)  //dp[i]代表凑成金额i需要的最少硬币数 也即一个最优子结构
	for i:=1; i<len(dp); i++ {
		dp[i] = amount+1        //初始化为amount+1 后续用于区分是否找不到解
	}
	for i:=1; i<len(dp); i++ {  //从amount == 1 开始dp
		for j:=0; j<len(coins); j++ {  //遍历coins列表找到dp[i-coins[i]]中的最小值 然后赋值给dp[i]
			if coins[j] <= i {  //若当前硬币面额大于amount 则说明不可由该硬币组成amount
				// 当兑换十元的时候，当前面额是五元，那么只需要找到五元的面额都有多少种方法即可。其余的同理
				dp[i] = min(dp[i], dp[i-coins[j]]+1)  // +1 表示添加了coins[i]这枚硬币
			}
		}
	}
	if dp[len(dp)-1] == amount+1 {  //若dp数组最后一个元素为amount+1说明其为发生变动 即不存在要求的硬币组合
		return -1
	}
	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
