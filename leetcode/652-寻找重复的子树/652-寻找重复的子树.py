# Definition for a binary tree node.
from re import sub
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 652-寻找重复的子树
class Solution:
    
    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        memo = {}
        res = []
        def traverse(root:Optional[TreeNode]):
            if root is None:
                return "#"
            left = traverse(root.left)
            right = traverse(root.right)
            sub_tree = "{} - {} - {}".format(left,right,root.val)
            if sub_tree in memo.keys():
                memo[sub_tree] = memo[sub_tree]+1
            else:
                memo[sub_tree] = 1
            if memo.get(sub_tree,0) == 2:
                res.append(root)
            
            return sub_tree
        traverse(root)
        return res