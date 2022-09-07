package main

import "sort"

// 90-子集 II
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	if len(nums) == 0 {
		return res
	}
	// 排序可以让相同的元素在一起
	sort.Ints(nums)
	backtrack(nums, &res, track, 0)
	return res
}

func backtrack(nums []int, res *[][]int, track []int, start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第⼀条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, res, track, i+1)
		track = track[:len(track)-1]
	}
}
