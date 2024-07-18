/* https://leetcode.com/problems/merge-two-binary-trees/description/

Вам даны два двоичных дерева root1 и root2.
Представьте, что когда вы накладываете одно из них на другое, некоторые узлы двух деревьев перекрываются, а другие нет.
Вам необходимо объединить два дерева в новое двоичное дерево. Правило слияния заключается в том, что если два узла
перекрываются, суммируйте значения узлов как новое значение объединенного узла. В противном случае узел NOT null будет
использоваться в качестве узла нового дерева.
Верните объединенное дерево.
Примечание. Процесс слияния должен начинаться с корневых узлов обоих деревьев.

Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]

Input: root1 = [1], root2 = [1,2]
Output: [2,2]
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
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {TreeNode}
 */
var mergeTrees = function(root1, root2) {
    if (!root1 && !root2) return null;
    if (!root1) return root2;
    if (!root2) return root1;

    root1.val += root2.val;
    root1.left = mergeTrees(root1.left, root2.left);
    root1.right = mergeTrees(root1.right, root2.right);

    return root1;
};

const example1 = [1,3,2,5];
const tree1 = buildTree(example1);
const example2 = [2,1,3,null,4,null,7];
const tree2 = buildTree(example2)
console.log(mergeTrees(tree1, tree2)); // [3,4,5,5,4,null,7]