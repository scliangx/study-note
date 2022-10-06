/**
 * Definition for a binary tree node.
 * */
function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @return {number}
 */

// 111-二叉树的最小深度
var minDepth = function (root) {
    if (!root) {
        return 0;
    }
    let queue = [root];
    let minVal = 1;
    while (queue.length > 0) {
        const size = queue.length;
        for (let i = 0; i < size; i++) {
            const node = queue.shift();
            if (node.left === null && node.right === null) {
                return minVal;
            }
            if (node.left !== null) {
                queue.push(node.left);
            }
            if (node.right != null) {
                queue.push(node.right);
            }
        }
        minVal++;
    }
    return minVal;
};