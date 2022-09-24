package main

// 567-字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	left,right,valid := 0,0,0
	for right < len(s2){
		c := s2[right]
		right++
		// 如果s1 存在c，则添加到滑动窗口Window中
		if need[c] > 0 {
			window[c]++
			// 其中一个字符满足，则valid增加
			if need[c] == window[c]{
				valid++
			}
		}
		for right - left >= len(s1){
			if valid == len(need){
				return true
			}
			d := s2[left]
			left++
			if need[d] > 0{
				if need[d] == window[d]{
					valid--
				}
				window[d]--
			}
		}

	}
	return false
}
