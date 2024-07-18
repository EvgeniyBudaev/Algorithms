/* https://leetcode.com/problems/path-sum/description/

Учитывая корень двоичного дерева и целое число targetSum, верните true, если дерево имеет путь от корня к листу, так что
суммирование всех значений по пути равняется targetSum.
Лист — это узел, не имеющий дочерних элементов.

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Объяснение: Показан путь от корня к листу с целевой суммой.
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
 * @param {number} targetSum
 * @return {boolean}
 */
var hasPathSum = function(root, targetSum) {
    // DFS
    if(!root) return false;
    if(!root.left && !root.right && targetSum === root.val) return true;
    return hasPathSum(root.left, targetSum - root.val)
        || hasPathSum(root.right, targetSum - root.val);
};

const example1 = [5,4,8,11,null,13,4,7,2,null,null,null,1];
const tree1 = buildTree(example1);
console.log(hasPathSum(tree1, 22)); // true