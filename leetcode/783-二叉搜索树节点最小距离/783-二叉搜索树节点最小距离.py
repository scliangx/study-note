# Definition for a binary tree node.
from re import L


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
# 783-二叉搜索树节点最小距离
class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        minVal = 100001
        pre = -1
        def traverse(root):
            if (root is None):
                return
            traverse(root.left)
            nonlocal pre,minVal
            if pre != -1 and root.val - pre < minVal:
                minVal = root.val - pre
            pre = root.val
            traverse(root.right)
        traverse(root)
        return minVal