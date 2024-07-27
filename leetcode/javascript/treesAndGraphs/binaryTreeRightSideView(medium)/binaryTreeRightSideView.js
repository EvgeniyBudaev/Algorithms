/* https://leetcode.com/problems/binary-tree-right-side-view/description/

Учитывая корень двоичного дерева, представьте, что вы стоите на его правой стороне, и верните значения узлов, которые
вы видите, упорядоченные сверху вниз.

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Input: root = [1,null,3]
Output: [1,3]

Input: root = []
Output: []
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
 * @return {number[]}
 */
var rightSideView = function(root) {
    if (!root) return [];

    const result = [];
    const queue = [root];

    while (queue.length) {
        result.push(queue[queue.length - 1].val);

        const len = queue.length;
        for (let i = 0; i < len; i++) {
            const node = queue.shift();

            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
        }
    }

    return result;
};

const example1 = [1,2,3,null,5,null,4];
const tree1 = buildTree(example1);
console.log(rightSideView(tree1)); // [1,3,4]
