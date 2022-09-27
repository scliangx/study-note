package main

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// 快慢指针，遇到奇数，原地交换位置
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]%2 != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}
