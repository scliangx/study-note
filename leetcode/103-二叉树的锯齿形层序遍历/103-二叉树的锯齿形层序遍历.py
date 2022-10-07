# Definition for a binary tree node.
from audioop import reverse
import re


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 103-二叉树的锯齿形层序遍历
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if (root is None):
            return []
        queue = [root]
        res = []
        while (len(queue) > 0):
            tmp = []
            size = len(queue)
            for i in range(size):
                node = queue.pop(0)
                tmp.append(node.val)
                if (not node.left is None):
                    queue.append(node.left)
                if (not node.right is None):
                    queue.append(node.right)
            res.append(tmp)
        for i in range(1,len(res),2):
            res[i].reverse()
        return res
