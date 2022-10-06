"""
# Definition for a Node.
"""
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

# 116-填充每个节点的下一个右侧节点指针
class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if root is None:
            return root
        q = [root]
        while len(q) > 0:
            size = len(q)
            tmp = []
            for i in range(size):
                node = q[0]
                q = q[1:]
                tmp.append(node)
                if not node.left is None:
                    q.append(node.left)
                if not node.right is None:
                    q.append(node.right)

            for i in range(len(tmp)-1):
                tmp[i].next = tmp[i+1]
        return root