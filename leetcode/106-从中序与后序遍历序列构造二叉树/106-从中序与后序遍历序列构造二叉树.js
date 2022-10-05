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
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */

//  106-从中序与后序遍历序列构造二叉树
var buildTree = function (inorder, postorder) {
    if (!postorder.length) {
        return null;
    }
    let top = postorder.pop()
    let index = inorder.indexOf(top) >= 0 ? inorder.indexOf(top) : 0;
    let root = new TreeNode(top);

    root.left = buildTree(inorder.slice(0, index + 1), postorder.slice(0, index));
    root.right = buildTree(inorder.slice(index + 1), postorder.slice(index));
    return root;
};