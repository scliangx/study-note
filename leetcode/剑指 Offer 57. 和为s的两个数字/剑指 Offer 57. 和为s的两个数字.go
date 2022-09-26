package main

// 剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
