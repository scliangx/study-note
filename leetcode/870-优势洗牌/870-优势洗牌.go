package main

import "sort"

type pair struct {
	index int
	value int
}

// 870-优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	sort.Ints(nums1)
	nums2Pair := make([]pair, len(nums2))
	// 利用新的结构存储nums2的值的顺序位置
	for i := 0; i < len(nums2); i++ {
		nums2Pair[i].value = nums2[i]
		nums2Pair[i].index = i
	}
	sort.Slice(nums2Pair, func(i, j int) bool {
		return nums2Pair[i].value < nums2Pair[j].value
	})
	cur, drop := 0, len(nums2)-1
	for i := range nums1 {
		// 如果nums1[i] 能跑过nums2[cur] 对应的位置，则直接跑
		if nums1[i] > nums2Pair[cur].value {
			res[nums2Pair[cur].index] = nums1[i]
			cur++
		} else {
			// 如果跑不过，直接找炮灰
			res[nums2Pair[drop].index] = nums1[i]
			drop--
		}
	}
	return res
}