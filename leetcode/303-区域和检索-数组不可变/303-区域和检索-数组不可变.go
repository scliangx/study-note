package main

type NumArray struct {
	res []int
}

// 构造一个前缀和数组
func Constructor(nums []int) NumArray {
	data := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		data[i] = data[i-1] + nums[i-1]
	}
	return NumArray{
		res: data,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.res[right+1] - this.res[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

