package main

import "fmt"

func main() {
	nums := []int{2, 3, 8, 4, 6, 12, 3, 4, 59, 1, 2, 7}

	fmt.Println(nums)
	fmt.Println(quickSort(nums, 0, len(nums)-1))
}

func quickSort(nums []int, low, high int) []int {
	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
	return nums
}

func partition(list []int, low, high int) int {
	tmp := list[low]
	for low < high {

		if low < high && tmp <= list[high] {
			high--
		}
		list[low] = list[high]

		if low < high && tmp >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = tmp
	return low
}
