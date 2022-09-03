package main

// 冒泡排序
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		exchange := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] { 
				nums[j], nums[j+1] = nums[j+1], nums[j] //交换
				exchange = true
			}
		}
		if !exchange {
			return nums
		}
	}
	return nums
}
