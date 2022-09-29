package main

// 238-除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0] = nums[0]
	suffix[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}
	res[0] = suffix[1]
	res[n-1] = prefix[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = prefix[i-1] * suffix[i+1]
	}
	return res
}
