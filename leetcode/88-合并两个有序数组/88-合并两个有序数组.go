package main

// 88-合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	p := len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}
	for j >= 0 {
		nums1[p] = nums2[j]
		p--
		j--
	}
}
