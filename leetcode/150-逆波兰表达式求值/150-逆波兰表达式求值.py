# 150-逆波兰表达式求值
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for i in tokens:
            if i == "+" or i == "-" or i == "*" or i == "/":
                n = len(stack)
                num1,num2 = stack[-1],stack[-2]
                stack = stack[:n-2]
                if i == "+":
                    stack.append(num1 + num2)
                elif i == "-":
                    stack.append(num2 - num1)
                elif i == "*":
                    stack.append(num1 * num2)
                elif i == "/":
                    stack.append(int(num2 / num1))
            else:
                stack.append(int(i))
            
        
        return stack[-1]