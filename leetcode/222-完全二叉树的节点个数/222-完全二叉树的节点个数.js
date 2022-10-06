/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */

// 222-完全二叉树的节点个数
var countNodes = function (root) {
    if (!root) { 
        return null;
    }
    let lDepth = 0,rDepth = 0;
    let l = root, r = root;
    while (l !== null) { 
        lDepth++;
        l = l.left;
    }
    
    while (r !== null) { 
        rDepth++;
        r = r.right;
    }
    if (lDepth === rDepth) { 
        return Number(Math.pow(2, lDepth) - 1);
    }
    return 1 + countNodes(root.left) + countNodes(root.right);
};