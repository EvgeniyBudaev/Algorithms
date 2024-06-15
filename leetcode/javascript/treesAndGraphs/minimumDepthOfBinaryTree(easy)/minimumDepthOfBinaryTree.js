/*
Учитывая двоичное дерево, найдите его минимальную глубину.
Минимальная глубина — это количество узлов на кратчайшем пути от корневого узла до ближайшего листового узла.
Примечание. Лист — это узел, не имеющий дочерних элементов.
*/

/*
Input: root = [3,9,20,null,null,15,7]
Output: 2
 */

/*
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5
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
var minDepth = function(root) {
    if (!root) return 0;
    if (root.left === null) return minDepth(root.right) + 1;
    if (root.right === null) return minDepth(root.left) + 1;
    return Math.min(minDepth(root.left), minDepth(root.right)) + 1;
};

const example1 = [3,9,20,null,null,15,7];
const tree1 = buildTree(example1);
const example2 = [2,null,3,null,4,null,5,null,6];
const tree2 = buildTree(example2);
console.log(minDepth(tree1)); // 2
console.log(minDepth(tree2)); // 5