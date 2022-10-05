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
 * @param {number[]} nums
 * @return {TreeNode}
 */

//  654-最大二叉树
var constructMaximumBinaryTree = function (nums) {
    const buildTree = function (start, end) {
        if (start > end) {
            return null;
        }
        let index = -1;
        let maxVal = -1;
        for (let i = start; i <= end; i++) {
            if (nums[i] > maxVal) {
                index = i;
                maxVal = nums[i];
            }
        }
        let root = new TreeNode(maxVal);
        root.left = buildTree(start, index - 1);
        root.right = buildTree(index + 1, end);

        return root;
    };
    return buildTree(0, nums.length - 1);
};
