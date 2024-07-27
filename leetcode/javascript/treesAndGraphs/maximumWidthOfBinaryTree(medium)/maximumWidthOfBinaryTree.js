/* https://leetcode.com/problems/maximum-width-of-binary-tree/description/

Учитывая корень двоичного дерева, верните максимальную ширину данного дерева.
Максимальная ширина дерева — это максимальная ширина среди всех уровней.
Ширина одного уровня определяется как длина между конечными узлами (крайний левый и крайний правый ненулевые узлы),
где нулевые узлы между конечными узлами, которые будут присутствовать в полном двоичном дереве, простирающемся до этого
уровня, равны также учитывается при расчете длины.
Гарантируется, что ответ будет находиться в диапазоне 32-битного целого числа со знаком.

Input: root = [1,3,2,5,3,null,9]
Output: 4
Объяснение: Максимальная ширина существует на третьем уровне с длиной 4 (5,3,null,9).
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
var widthOfBinaryTree = function(root) {
    const minPos = [0];
    let maxWidth = 0;

    callDFS(root, 0, 0);
    return maxWidth;

    function callDFS(node, level, pos) {
        if(!node) return;
        if(minPos[level] === undefined) minPos.push(pos);

        const diff = pos - minPos[level];
        maxWidth = Math.max(maxWidth, diff + 1);

        callDFS(node.left, level + 1, diff *2 );
        callDFS(node.right, level + 1, diff * 2 + 1);
    }
};

const example1 = [1,3,2,5,3,null,9];
const tree1 = buildTree(example1);
console.log(widthOfBinaryTree(tree1)); // 4