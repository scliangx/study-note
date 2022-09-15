package main

// 875-爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	// 最低每次吃一个,最多每次吃香蕉的最大限制数量个
	left,right := 1,1000000000 + 1
	for left < right{
		mid := left + (right - left) / 2
		if f(piles,mid) <= h{
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}

func f(piles []int,x int )int {
	hours := 0
	for i:=0;i<len(piles);i++{
		// 当前堆有piles[i]个香蕉，速度为x，则吃当前一堆香蕉时间为: piles[i] / x
		hours += piles[i] / x
		// 如果piles[i] % x == 0 表示可整数小时吃完，如果不等于0,则还额外需要一个小时
		if piles[i] % x > 0 {
			hours++
		}
	}
	return hours
}
