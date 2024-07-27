/* https://leetcode.com/problems/validate-binary-search-tree/description/

Учитывая корень двоичного дерева, определите, является ли оно допустимым двоичным деревом поиска (BST).
Действительный BST определяется следующим образом:
Левый поддерево узла содержит только узлы с ключами меньшими, чем ключ узла.
Правое поддерево узла содержит только узлы с ключами, превышающими ключ узла.
И левое, и правое поддеревья также должны быть бинарными деревьями поиска.

Input: root = [2,1,3]
Output: true

Input: root = [5,1,4,null,null,3,6]
Output: false
Объяснение: Значение корневого узла равно 5, а значение его правого дочернего узла равно 4.
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
var isValidBST = function(root, min = null, max = null) {
    if (!root) return true;
    if (min !== null && root.val <= min) return false;
    if (max !== null && root.val >= max) return false;

    const rLeft = isValidBST(root.left, min, root.val);
    const rRight = isValidBST(root.right, root.val, max);

    return rLeft && rRight;
};

const example1 = [2,1,3];
const tree1 = buildTree(example1);
console.log(isValidBST(tree1)); // true