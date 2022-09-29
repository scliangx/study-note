# 71-简化路径

class Solution:
    def simplifyPath(self, path: str) -> str:
        paths = path.split("/")
        stack = []
        for i in paths:
            if i == "..":
                if len(stack) > 0:
                    stack = stack[:len(stack)-1]
            elif i !=""  and i != ".":
                stack.append(i)

        return "/" + "/".join(stack)
             

        