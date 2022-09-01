package main

// 34-在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	// 找出最左边的target的位置
	for left <= right {
		mid := right - (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	// 计算最右边的位置
	cur := left
	for i := left + 1; i < len(nums); i++ {
		if nums[i] == target {
			cur++
		}
	}
	return []int{left, cur}
}
