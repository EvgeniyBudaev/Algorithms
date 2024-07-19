/* https://leetcode.com/problems/binary-tree-paths/description/

Учитывая корень двоичного дерева, верните все пути от корня к листу в любом порядке.
Лист — это узел, не имеющий дочерних элементов.

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Input: root = [1]
Output: ["1"]
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
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    const paths = [];

    function findPaths(node, path) {
        if (!node) return;

        path.push(node.val);

        if (!node.left && !node.right) {
            paths.push(path.join('->'));
        } else {
            findPaths(node.left, [...path]);
            findPaths(node.right, [...path]);
        }
    }

    findPaths(root, []);
    return paths;
};

const example1 = [1,2,3,null,5];
const tree1 = buildTree(example1);
console.log(binaryTreePaths (tree1)); // ["1->2->5","1->3"]