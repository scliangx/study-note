package main

import "fmt"

// 插入排序
func main() {
	nums := [5]int{10, 54, 36, 89, -1}
	fmt.Println("nums = ", nums)
	InsertSort(&nums)
	fmt.Println("nums = ", nums)
}

func InsertSort(arr *[5]int) {
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
