package main

//剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	// 查找到了最低有一个
	count := 1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 没有查找到或者索引越界
	if left >= len(nums) || nums[left] != target {
		return 0
	}
	// 因为是有序的，所以看后边有多少个即可
	for i := left + 1; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}
