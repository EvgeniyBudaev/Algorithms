/* https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

Учитывая корень двоичного дерева поиска и целое число k, верните k-е наименьшее значение (с индексом 1) из всех значений
узлов в дереве.

Input: root = [3,1,4,null,2], k = 1
Output: 1

Input: root = [5,3,6,2,4,null,null,1], k = 3
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
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function(root, k) {
    let n = 0;
    let stack = [];
    let current = root;
    while (current || stack.length > 0) {
        while (current) {
            stack.push(current);
            current = current.left;
        }
        current = stack.pop();
        n += 1;
        if (n === k) return current.val;
        current = current.right;
    }
};

const example1 = [3,1,4,null,2];
const tree1 = buildTree(example1);
console.log(kthSmallest(tree1, 1)); // 1