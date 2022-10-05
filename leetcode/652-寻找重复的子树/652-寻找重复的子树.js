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
 * @return {TreeNode[]}
 */

// 652-寻找重复的子树
var findDuplicateSubtrees = function (root) {
    let memo = new Map();
    let res = new Array();
    const traverse = function (root) {
        if (!root) {
            return "#";
        }
        const left = traverse(root.left);
        const right = traverse(root.right);
        const subTree = `${left},${right},${root.val}`;
        const count = memo.get(subTree);
        memo.set(subTree, (count || 0) + 1);
        if (memo.get(subTree) == 2) res.push(root);
        return subTree;
    };
    traverse(root);
    return res;
};