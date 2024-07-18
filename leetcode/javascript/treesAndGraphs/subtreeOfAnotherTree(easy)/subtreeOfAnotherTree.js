/* https://leetcode.com/problems/subtree-of-another-tree/description/
solution https://leetcode.com/problems/subtree-of-another-tree/solutions/1479754/javascript-98-solutions-for-bfs-and-dfs/

Учитывая корни двух двоичных деревьев root и subRoot, верните true, если существует корневое поддерево с той же
структурой и значениями узлов subRoot, и false в противном случае.
Поддерево двоичного дерева — это дерево, состоящее из узла дерева и всех его потомков. Дерево-дерево также можно
рассматривать как поддерево самого себя.

Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
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
 * @param {TreeNode} subRoot
 * @return {boolean}
 */
var isSubtree = function(root, subRoot) {
    const areEqual = (node1, node2) => {
        if (!node1 || !node2) return !node1 && !node2;
        if (node1.val !== node2.val) return false;
        return areEqual(node1.left, node2.left) && areEqual(node1.right, node2.right);
    }
    const dfs = (node) => {
        if (!node) return false;
        if (areEqual(node, subRoot)) return true;
        return dfs(node.left) || dfs(node.right);
    }
    return dfs(root);
};

const example1 = [3,4,5,1,2];
const tree1 = buildTree(example1);
const example2 = [4,1,2];
const tree2 = buildTree(example2);
console.log(isSubtree(tree1, tree2)); // true