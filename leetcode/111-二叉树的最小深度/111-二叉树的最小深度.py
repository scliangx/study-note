# Definition for a binary tree node.
from operator import le


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 111-二叉树的最小深度
class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if (root is None):
            return 0
        queue = [root]
        minVal = 1
        while (len(queue) > 0 ):
            size = len(queue)
            for i in range(size):
                node = queue[0]
                queue = queue[1:]
                if (node.left is None and node.right is None):
                    return minVal
                if (not node.left is None):
                    queue.append(node.left)
                if (not node.right is None):
                    queue.append(node.right)
            minVal += 1
        return minVal