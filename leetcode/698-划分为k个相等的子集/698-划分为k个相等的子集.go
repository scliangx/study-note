package main

// 698-划分为k个相等的子集
// 从数字的角度看
// 运行超时
/*
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	target := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target = sum / k
	bucket := make([]int, k)
	return backtrack(nums, 0, target, bucket)
}

func backtrack(nums []int, index int, target int, bucket []int) bool {
	// 分成k个子集，bucket中是每个子集的和，检查每个子集的和是否等于target
	if index == len(nums) {
		for i := 0; i < len(bucket); i++ {
			if bucket[i] == target {
				return true
			}
		}
		return false
	}
	// 讲所有的数字依次放到桶的某一个位置，并且所有位置的和需要相等
	for i := 0; i < len(bucket); i++ {
		// 剪枝
		if bucket[i]+nums[index] > target {
			continue
		}
		bucket[i] += nums[index]
		if backtrack(nums, index+1, target, bucket) {
			return true
		}
		bucket[i] -= nums[index]
	}
	return false
}
*/


// 698-划分为k个相等的子集
func canPartitionKSubsets(nums []int, k int) bool {
	// 边界条件检查
	if len(nums) == 0 {
		return false
	}
	sum := 0
	target := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target = sum / k
	used := make([]bool, len(nums))
	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	return backtrack(k, 0, nums, 0, used, target)
}

func backtrack(k int, bucket int, nums []int, start int, used []bool, target int) bool {
	// 所有桶都被装满了，⽽且 nums ⼀定全部⽤完了
	if k == 0 {
		return true
	}
	if bucket == target {
		// 装满了当前桶，递归穷举下⼀个桶的选择
		// 让下⼀个桶从 nums[0] 开始选数字
		return backtrack(k-1, 0, nums, 0, used, target)
	}
	// 从 start 开始向后探查有效的 nums[i] 装⼊当前桶
	for i := start; i < len(nums); i++ {
		// 当前桶装不下 nums[i]，跳过
		if used[i] || bucket+nums[i] > target {
			continue
		}
		used[i] = true
		// 做选择，将 nums[i] 装⼊当前桶中
		bucket += nums[i]
		// 递归穷举下⼀个数字是否装⼊当前桶
		if backtrack(k, bucket, nums, i+1, used, target) {
			return true
		}
		// 撤销选择
		used[i] = false
		bucket -= nums[i]
	}
	return false
}
