package main

// 45-跳跃游戏 II
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	end := 0
	count := 0
	maxP := 0
	for i := 0; i < len(nums)-1; i++ {
		maxP = max(maxP, i+nums[i])
		if i == end {
			end = maxP
			count++
		}
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
