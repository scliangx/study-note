package main

// 921-使括号有效的最少添加
func minAddToMakeValid(s string) int {
	res := 0
	need := 0
	for i := range s {
		// 需要多少个右括号
		if s[i] == '(' {
			need++
		}
		// 如果是右括号，则右括号的需求数量减一
		if s[i] == ')' {
			need--
			// 右括号的需求数量减少到-1，那么久必须有一次添加操作
			if need == -1 {
				need = 0
				res++
			}
		}
	}
	return res + need
}
