

# 965-单值二叉树
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    
    def isUnivalTree(self, root: Optional[TreeNode]) -> bool:  
        if root is None:
            return True
        if root.left:
            if root.val != root.left.val or not self.isUnivalTree(root.left):
                return False
        if root.right:
            if root.val != root.right.val or not self.isUnivalTree(root.right):
                return False

        return True
        