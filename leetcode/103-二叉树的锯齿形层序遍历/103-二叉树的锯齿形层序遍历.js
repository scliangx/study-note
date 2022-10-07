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

// 103-二叉树的锯齿形层序遍历
var zigzagLevelOrder = function (root) {
    if (!root) { 
        return [];
    }
    let res = [];
    const queue = [root];
    while (queue.length > 0) { 
        const tmp = [];
        const size = queue.length;
        for (let i = 0; i < size; i++) { 
            const node = queue.shift();
            tmp.push(node.val);
            if (node.left !== null) { 
                queue.push(node.left);
            }
            if (node.right !== null) { 
                queue.push(node.right);
            }
        }
        res.push(tmp);
    }
    for (let i = 1; i < res.length; i += 2) { 
        res[i].reverse();
    }
    return res;
};