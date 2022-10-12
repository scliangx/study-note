package main

import "strconv"

// 443-压缩字符串
func compress(chars []byte) int {
	res := 0
	cur := 0
	for r := 0; r < len(chars); r++ {
		// 走到最后一个或者找到不一样的了，统计前边可以压缩的
		if r+1 == len(chars) || chars[r+1] != chars[cur] {
			chars[res] = chars[cur]
			res++
			// 统计相同的部分
			if r > cur {
				for _, v := range strconv.Itoa(r + 1 - cur) {
					chars[res] = byte(v)
					res++
				}
			}
			cur = r + 1
		}
	}
	return res
}
