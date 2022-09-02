package main

import "math/rand"

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

// 380-O(1) 时间插入、删除和获取随机元素

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	// 从哈希表中获得 val 的下标 index；
	index, ok := rs.indices[val]
	if !ok {
		return false
	}

	lastIndex := len(rs.nums) - 1
	// 将变长数组的最后一个元素last 移动到下标 index 处
	rs.nums[index] = rs.nums[lastIndex]
	// 在哈希表中将last 的下标更新为index；
	rs.indices[rs.nums[index]] = index
	// 在变长数组中删除最后一个元素
	rs.nums = rs.nums[:lastIndex]
	// 在哈希表中删除 val；
	delete(rs.indices, val)
	//  返回 true
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
