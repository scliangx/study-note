package main

import "sort"

// 40-组合总和 II
// 额外记录路劲和，base case更改即可
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	backtrack(candidates, 0, 0, track, target, &res)
	return res
}

func backtrack(candidates []int, start int, sumVal int, track []int, target int, res *[][]int) {
	if sumVal == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	// 提前结束
	if sumVal > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		sumVal += candidates[i]
		track = append(track, candidates[i])
		backtrack(candidates, i+1, sumVal, track, target, res)
		track = track[:len(track)-1]
		sumVal -= candidates[i]
	}
}
