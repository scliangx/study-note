/**
 * Definition for a binary tree node.
 * 
 */
function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @return {number[][]}
 */

// 107-二叉树的层序遍历 II
var levelOrderBottom = function (root) {
    if (!root) { 
        return [];
    }
    const queue = [root];
    const res = [];
    while (queue.length > 0) { 
        let tmp = [];
        const size = queue.length;
        for (let i = 0; i < size; i++) { 
            node = queue.shift()
            tmp.push(node.val);
            if (node.left !== null) { 
                queue.push(node.left);
            }
            if (node.right !== null) { 
                queue.push(node.right)
            }
        }
        res.push(tmp);
    }
    res.reverse();
    return res;
};