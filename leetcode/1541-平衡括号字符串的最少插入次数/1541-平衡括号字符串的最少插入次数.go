package main

// 
func minInsertions(s string) int {
	res := 0
	need := 0
	for i := range s {
		if s[i] == '(' {
			need += 2
			// 另外，当遇到左括号时，若对右括号的需求量为奇数，需要插⼊ 1 个右括号。
			// 因为⼀个左括号需要两个右括号，右括号的需求必须是偶数
			if need%2 == 1 {
				// 插⼊⼀个右括号
				res++
				// 对右括号的需求减⼀
				need--
			}
		}
		if s[i] == ')' {
			need--
			if need == -1 {
				need = 1
				res++
			}
		}
	}
	return res + need
}
