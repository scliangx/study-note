package main

// 496-下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// 当前栈不为空；如果是挨个子直接删掉
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 如果栈不为空 找到
		if len(stack) > 0 {
			set[num] = stack[len(stack)-1]
		} else {
			set[num] = -1
		}
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = set[v]
	}
	return res
}
