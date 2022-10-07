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

// 1373-二叉搜索子树的最大键值和
var maxSumBST = function (root) {
    if (!root) {
        return 0;
    }
    let res = 0;
    var traverse = function (root) {
        // 定义⼀个 traverse 函数，traverse(root) 返回⼀个⼤⼩为 4 的 int 数组，我们暂且称它为 res，其 中：
        //     res[0] 记录以 root 为根的⼆叉树是否是 BST，若为 1 则说明是 BST，若为 0 则说明不是 BST；
        //     res[1] 记录以 root 为根的⼆叉树所有节点中的最⼩值；
        //     res[2] 记录以 root 为根的⼆叉树所有节点中的最⼤值；
        //     res[3] 记录以 root 为根的⼆叉树所有节点值之和。
        if (root === null) {
            return {
                isBST: true,
                minVal: Infinity,
                maxVal: -Infinity,
                sumVal: 0
            };
        }
        const left = traverse(root.left);
        const right = traverse(root.right);
        if (left.isBST && right.isBST && left.maxVal < root.val && root.val < right.minVal) {
            let sum = left.sumVal + right.sumVal + root.val
            res = Math.max(sum, res);
            return {
                isBST: true,
                minVal: Math.min(root.val, left.minVal),
                maxVal: Math.max(root.val, right.maxVal),
                sumVal: sum,
            };
        } else {
            return { isBST: false };
        }

    };
    traverse(root);
    return res;
};