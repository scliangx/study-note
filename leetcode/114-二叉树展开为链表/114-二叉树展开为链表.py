
# Definition for a binary tree node.
from math import fabs
from operator import le


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
# 114-二叉树展开为链表
class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if root is None:
            return root
        
        self.flatten(root.left)
        self.flatten(root.right)
        
        left = root.left
        right = root.right
        root.right = left
        root.left = None

        cur = root
        while not cur.right is None:
            cur = cur.right
        cur.right = right
