package main

import "strconv"

// 150-逆波兰表达式求值
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return -1
	}
	stack := []int{}
	for _, v := range tokens {
		if v == "*" || v == "/" || v == "-" {
			n := len(stack)
			num1, num2 := stack[n-1], stack[n-2]
			stack = stack[:n-2]
			switch v {
			case "*":
				val := num1 * num2
				stack = append(stack, val)
			case "/":
				val := num2 / num1
				stack = append(stack, val)
			case "-":
				val := num2 - num1
				stack = append(stack, val)
			}
		} else {
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		}
	}
	return stack[len(stack)-1]
}
