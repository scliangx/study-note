package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 8, 4, 6, 12, 3, 4, 59, 1, 2, 7}

	fmt.Println(nums)
	fmt.Println(selectionSort(nums))
}

// selectionSort 选择排序
func selectionSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if nums[i] >= nums[min] {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}
