# Definition for a binary tree node.
from re import S


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 226-翻转二叉树
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        queue = [root]
        while len(queue) > 0:
            size = len(queue)
            for i in range(size):
                node = queue[0]
                queue = queue[1:]
                node.left,node.right = node.right,node.left
                if not node.left is None:
                    queue.append(node.left)
                if not node.rigt is None:
                    queue.append(node.right)
        return root