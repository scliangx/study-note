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
 * @return {void} Do not return anything, modify root in-place instead.
 */

// 114-二叉树展开为链表
var flatten = function (root) {
    if (!root) { 
        return;
    }
    flatten(root.left);
    flatten(root.right);

    const left = root.left;
    const right = root.right;
    root.right = left;
    root.left = null;

    const cur = root;
    while (root.right) {
        cur = cur.right;
    }
    cur.right = right;
};