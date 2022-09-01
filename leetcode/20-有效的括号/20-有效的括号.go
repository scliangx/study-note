package main

// 20-有效的括号
func isValid(s string) bool {
	stack := []byte{}
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) != 0 && leftOf(s[i]) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func leftOf(c byte) byte {
	if c == '}' {
		return '{'
	} else if c == ']' {
		return '['
	} else {
		return '('
	}
}
