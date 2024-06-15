/*
Учитывая корень двоичного дерева, верните сумму значений его самых глубоких листьев.
*/

/*
Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15
 */

/*
Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 19
 */

class TreeNode {
    constructor(val, left = null, right = null) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

function buildTree(values) {
    const root = new TreeNode(values.shift());
    const queue = [root];
    while (queue.length > 0 && values.length > 0) {
        const curNode = queue.shift();
        const leftVal = values.shift();
        const rightVal = values.shift();
        if (leftVal !== null) {
            curNode.left = new TreeNode(leftVal);
            queue.push(curNode.left);
        }
        if (rightVal !== null) {
            curNode.right = new TreeNode(rightVal);
            queue.push(curNode.right);
        }
    }
    return root;
}

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
var deepestLeavesSum = function(root) {
    let sums = [];

    const dfs = (node, lvl) => {
        if (lvl === sums.length) sums[lvl] = node.val;
        else sums[lvl] += node.val;
        if (node.left) dfs(node.left, lvl+1);
        if (node.right) dfs(node.right, lvl+1);
    }

    dfs(root, 0);
    return sums[sums.length - 1];
};

const example1 = [1,2,3,4,5,null,6,7,null,null,null,null,8];
const tree1 = buildTree(example1);
console.log(deepestLeavesSum(tree1)); // 15