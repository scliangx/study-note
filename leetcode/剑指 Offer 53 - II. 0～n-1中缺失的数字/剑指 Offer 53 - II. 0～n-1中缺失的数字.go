package main

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// 0 ~ n-1 那么，根据nums[mid] <|> mid可以判断出缺失元素的位置
func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		// 说明缺失的数字在数组中间的左侧
		if nums[mid] > mid {
			right = mid - 1
		} else {
			// 右侧
			left = mid + 1
		}
	}
	return left
}
