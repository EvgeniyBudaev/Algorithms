/*
Учитывая корень двоичного дерева, верните обход зигзагообразного порядка значений его узлов.
(т. е. слева направо, затем справа налево для следующего уровня и попеременно)
*/

/*
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
 */

/*
Input: root = [1]
Output: [[1]]
 */

/*
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
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    let res = [];

    const dfs = (node, h) => {
        if (!node) return;
        if (!res[h]) res[h] = [];

        if (h % 2 === 0) { // четная высота
            res[h].push(node.val);
        } else { // нечетная высота
            res[h].unshift(node.val);
        }

        dfs(node.left, h + 1);
        dfs(node.right, h + 1);
    }

    dfs(root, 0);
    return res;
};

const example1 = [3,9,20,null,null,15,7];
const tree1 = buildTree(example1);
console.log(zigzagLevelOrder(tree1)); // [[3],[20,9],[15,7]]