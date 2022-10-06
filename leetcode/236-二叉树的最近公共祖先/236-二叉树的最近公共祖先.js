/**
 * Definition for a binary tree node.
 */
function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */

// 236-二叉树的最近公共祖先
var lowestCommonAncestor = function (root, p, q) {
    if (!root) {
        return null;
    }
    if (root === p || root === q) {
        return root;
    }
    const left = lowestCommonAncestor(root.left, p, q)
    const right = lowestCommonAncestor(root.right, p, q)
    if (!left && !right) {
        return null;
    }
    if (left && right) {
        return root;
    }
    return left === null ? right : left
};