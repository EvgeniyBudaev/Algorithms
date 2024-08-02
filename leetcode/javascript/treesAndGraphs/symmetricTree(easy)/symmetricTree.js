/* https://leetcode.com/problems/symmetric-tree/description/

Учитывая корень двоичного дерева, проверьте, является ли оно зеркалом самого себя (т. е. симметрично относительно своего
центра).

Input: root = [1,2,2,3,4,4,3]
Output: true

Input: root = [1,2,2,null,3,null,3]
Output: false
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
 * @return {boolean}
 */
var isSymmetric = function(root) {
    if (!root) return true;

    function helper(root1, root2) {
        if (root1 === null && root2 === null) return true;
        if (root1 === null || root2 === null) return false;
        if (root1.val !== root2.val) return false;
        return helper(root1.left, root2.right) && helper(root1.right, root2.left);
    }

    return helper(root.left, root.right);
};

const example1 = [1,2,2,3,4,4,3];
const tree1 = buildTree(example1);
console.log(isSymmetric(tree1)); // true