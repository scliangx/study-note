package main

// 1352-最后 K 个数的乘积
type ProductOfNumbers struct {
	Val []int
}

func Constructor() ProductOfNumbers {
	l := []int{1}
	return ProductOfNumbers{Val: l}

}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.Val = []int{1}
		return
	}
	size := len(this.Val)
	// 前缀积数组中每个元素
	tmpArr := this.Val
	tmpVal := this.GetProduct(size - 1)
	tmpArr = append(tmpArr, tmpVal*num)
	this.Val = tmpArr
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.Val)
	if k > n-1 {
		// 不⾜ k 个元素，是因为最后 k 个元素存在 0
		return 0
	}
	// 计算最后 k 个元素积
	return (this.Val)[n-1] / (this.Val)[n-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
