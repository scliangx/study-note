package main

// 55-跳跃游戏
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
