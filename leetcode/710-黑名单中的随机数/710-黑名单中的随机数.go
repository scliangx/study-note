package main

import (
	"fmt"
	"math/rand"
)

// 710-黑名单中的随机数
type Solution struct {
	size    int
	mapping map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	mapping := map[int]int{}
	sz := n - len(blacklist)
	for _, v := range blacklist {
		mapping[v] = 666
	}

	last := n - 1
	check := func(m map[int]int, key int) bool {
		_, ok := m[key]
		return ok
	}
	for _, v := range blacklist {
		if v >= sz {
			continue
		}
		for check(mapping, last) {
			last--
		}
		mapping[v] = last
		last--
	}
	return Solution{mapping: mapping, size: sz}
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.size)
	fmt.Println("index: ", index)
	if v, ok := this.mapping[index]; ok {
		return v
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
