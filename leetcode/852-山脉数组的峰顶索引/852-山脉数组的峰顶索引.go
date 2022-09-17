package main

// 852-山脉数组的峰顶索引
func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	// 二分查找，两端都闭
	left, right := 0, len(arr)-1
	for left < right {
		// 取中间索引，要不在峰值左边，要不在峰值右边
		mid := left + (right-left)/2
		// 如果右边是递减,right 收缩
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
