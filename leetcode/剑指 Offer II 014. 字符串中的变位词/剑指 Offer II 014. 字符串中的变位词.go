package main

// 剑指 Offer II 014. 字符串中的变位词
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 {
		return true
	}
	need, window := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0
	valid := 0
	for i := range s1 {
		need[s1[i]]++
	}
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] > 0 {
			window[c]++
			// s1 中满足s2的字符的数量
			if need[c] == window[c] {
				valid++
			}
		}
		// 判断左侧窗⼝是否要收缩
		for right-left >= len(s1) {
			// 判断是否找到了合法的⼦串
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if need[d] > 0 {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的⼦串
	return false
}
