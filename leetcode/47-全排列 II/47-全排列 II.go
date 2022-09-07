package main

import "sort"

// 47-全排列 II
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	visited := make([]bool, len(nums))
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	backtrack(nums, track, visited, &res)
	return res
}

func backtrack(nums []int, track []int, visited []bool, res *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// 维持了相同值得一个相对的顺序
		// !visited[i-1] 只有在前一个值被选择的情况下，后一个值才会被选择
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		track = append(track, nums[i])
		backtrack(nums, track, visited, res)
		track = track[:len(track)-1]
		visited[i] = false
	}
}
