/* https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

Учитывая корень двоичного дерева, верните его максимальную глубину.
Максимальная глубина двоичного дерева — это количество узлов на самом длинном пути от корневого узла до самого дальнего
листового узла.

Input: root = [3,9,20,null,null,15,7]
Output: 3
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
var maxDepth = function(root) {
    if (!root) return 0;
    if (root.left === null) return maxDepth(root.right) + 1;
    if (root.right === null) return maxDepth(root.left) + 1;
    return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
};

const example1 = [3,9,20,null,null,15,7];
const tree1 = buildTree(example1);
console.log(maxDepth(tree1)); // 3