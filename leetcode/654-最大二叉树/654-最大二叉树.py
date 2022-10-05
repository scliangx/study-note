# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 654-最大二叉树
class Solution:
    def constructMaximumBinaryTree(self, nums: List[int]) -> Optional[TreeNode]:
        if len(nums) == 0:
             return None
        def build_tree(nums,start,end):
            if start > end:
                return None
            index,max_val = 0,-sys.maxsize - 1
            for i in range(start,end+1):
                if nums[i] > max_val:
                    index = i
                    max_val = nums[i]
            root = TreeNode(max_val)
            root.left = build_tree(nums,start,index-1)
            root.right = build_tree(nums,index+1,end)
            return root
        
        return build_tree(nums,0,len(nums)-1)