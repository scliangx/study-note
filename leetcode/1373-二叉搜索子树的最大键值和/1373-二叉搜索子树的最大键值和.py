from operator import le
import re
import sys

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# 1373-二叉搜索子树的最大键值和
class Solution:
    def maxSumBST(self, root: Optional[TreeNode]) -> int:
        max_value = 0
        def traverse(root):
            """
            定义⼀个 traverse 函数，traverse(root) 返回⼀个⼤⼩为 4 的 int 数组，我们暂且称它为 res，其 中：
            res[0] 记录以 root 为根的⼆叉树是否是 BST，若为 1 则说明是 BST，若为 0 则说明不是 BST；
            res[1] 记录以 root 为根的⼆叉树所有节点中的最⼩值；
            res[2] 记录以 root 为根的⼆叉树所有节点中的最⼤值；
            res[3] 记录以 root 为根的⼆叉树所有节点值之和。
            """
            if (root is None):
                return [1,sys.maxsize,-sys.maxsize-1,0]
            
            left = traverse(root.left)
            right = traverse(root.right)
            res = [0,0,0,0]
            if (left[0] ==1 and right[0] == 1 and left[2] < root.val and root.val <right[1]):
                res[0] = 1
                res[1] = min(root.val,left[1])
                res[2] = max(root.val,right[2])
                res[3] = left[3] + right[3] + root.val
                nonlocal max_value
                max_value = max(res[3],max_value)
            return res
        traverse(root)
        return max_value