/*
Учитывая корень двоичного дерева, верните длину диаметра дерева.
Диаметр двоичного дерева — это длина самого длинного пути между любыми двумя узлами дерева.
Этот путь может проходить через корень, а может и не проходить.
Длина пути между двумя узлами представлена количеством ребер между ними.
*/

/*
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
 */

/*
Input: root = [1,2]
Output: 1
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

var diameterOfBinaryTree = function(root) {
    let diameter = 0;

    function dfs(node) {
        if (!node) return 0;
        let left = dfs(node.left);
        let right = dfs(node.right);
        diameter = Math.max(diameter, left + right);
        return Math.max(left, right) + 1;
    }

    dfs(root);
    return diameter;
};

const example1 = [1,2,3,4,5];
const tree1 = buildTree(example1);
const example2 = [1,2];
const tree2 = buildTree(example2);
console.log(diameterOfBinaryTree(tree1)); // 3
console.log(diameterOfBinaryTree(tree2)); // 1