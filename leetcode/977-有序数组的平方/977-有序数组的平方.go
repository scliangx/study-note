package main

// 977-有序数组的平方
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	i, j := 0, len(nums)-1
	p := len(nums) - 1
	res := make([]int, len(nums))
	for i <= j {
		if abs(nums[i]) >= abs(nums[j]) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
		p--
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return 0 - x
}
