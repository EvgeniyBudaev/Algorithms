/* https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

Учитывая двоичное дерево поиска (BST), найдите узел наименьшего общего предка (LCA) из двух заданных узлов в BST.
Согласно определению LCA в Википедии: «Наименьший общий предок определяется между двумя узлами p и q как самый нижний
узел в T, который имеет как p, так и q в качестве потомков (где мы позволяем узлу быть потомком самого себя). »

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Пояснение: LCA узлов 2 и 8 равен 6.
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
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
    if (root === null) return root;

    if (root.val > p.val && root.val > q.val) lowestCommonAncestor(root.left, p, q);
    else if (root.val < p.val && root.val < q.val) lowestCommonAncestor(root.right, p, q);

    return root;
};

const example1 = [6,2,8,0,4,7,9,null,null,3,5];
const tree1 = buildTree(example1);
const tree2 = buildTree(2);
const tree3 = buildTree(8);
console.log(lowestCommonAncestor(tree1, tree2, tree3)); // 6