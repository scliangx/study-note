import math

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

#  222-完全二叉树的节点个数
class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        l,r = root,root
        l_depth,r_depth = 0,0
        while (not l is None):
            l_depth += 1
            l = l.left
        while (not r is None):
            r_depth += 1
            r = r.right
        if l_depth == r_depth:
            return int(math.pow(2,l_depth)-1)
        return 1 + self.countNodes(root.left) + self.countNodes(root.right)
