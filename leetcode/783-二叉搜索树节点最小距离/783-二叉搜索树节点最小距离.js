/**
 * Definition for a binary tree node.
 */
function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @return {number}
 */

// 783-二叉搜索树节点最小距离
var minDiffInBST = function (root) {
    if (!root) {
        return 0;
    }
    let minVal = Infinity;
    let pre = -1;
    const traverse = function (root) {
        if (!root) {
            return;
        }
        traverse(root.left);
        if (pre !== -1 && root.val - pre < minVal) {
            minVal = root.val - pre;
        }
        pre = root.val;
        traverse(root.right);
    };
    traverse(root);
    return minVal;
};