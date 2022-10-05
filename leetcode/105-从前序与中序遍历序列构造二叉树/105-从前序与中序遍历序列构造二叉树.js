/**
 * Definition for a binary tree node.
 *  */
function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */

//  105-从前序与中序遍历序列构造二叉树
var buildTree = function (preorder, inorder) {
    if (!inorder.length) {
        return null;
    }
    let root = new TreeNode(preorder[0]);
    let index = inorder.indexOf(preorder[0]) >= 0 ? inorder.indexOf(preorder[0]) : 0;

    root.left = buildTree(preorder.slice(1, index + 1), inorder.slice(0, index));
    root.right = buildTree(preorder.slice(index + 1), inorder.slice(index + 1));
    return root;
};