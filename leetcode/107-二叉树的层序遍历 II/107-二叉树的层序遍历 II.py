# Definition for a binary tree node.
from operator import le


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 107-二叉树的层序遍历 II
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        res = []
        queue = [root]
        while (len(queue) > 0):
            tmp = []
            size = len(queue)
            for i in range(size):
                node = queue[0]
                queue = queue[1:]
                tmp.append(node.val)
                if (not node.left is None):
                    queue.append(node.left)
                if (not node.right is None):
                    queue.append(node.right)
            res.append(tmp)
        l,r = 0,len(res)-1
        while l < r:
            res[l],res[r] = res[r],res[l]
            l += 1
            r -= 1
        return res