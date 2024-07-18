/* https://leetcode.com/problems/invert-binary-tree/description/
solution https://leetcode.com/problems/invert-binary-tree/solutions/399221/clean-javascript-iterative-dfs-bfs-solutions/

Учитывая корень двоичного дерева, инвертируйте дерево и верните его корень.

Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
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
 * @return {TreeNode}
 */
var invertTree = function(root) {
    // Recursion
    helper(root);
    return root;
};

function helper(root){
    if(root == null) return;

    let temp = root.left;
    root.left = root.right;
    root.right = temp;

    helper(root.left);
    helper(root.right);
}

var invertTreeWithDFS = function(root) {
    // DFS
    const stack = [root];
    while (stack.length) {
        const n = stack.pop();
        if (n != null) {
            [n.left, n.right] = [n.right, n.left];
            stack.push(n.left, n.right);
        }
    }
    return root;
};

function invertTreeWithBFS(root) {
    // BFS
    const queue = [root];
    while (queue.length) {
        const n = queue.shift();
        if (n != null) {
            [n.left, n.right] = [n.right, n.left];
            queue.push(n.left, n.right);
        }
    }
    return root;
}

const example1 = [4,2,7,1,3,6,9];
const tree1 = buildTree(example1);
console.log(invertTree(tree1)); // [4,7,2,9,6,3,1]