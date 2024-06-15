/*
Учитывая непустое двоичное дерево поиска и целевое значение, найдите в BST значение, ближайшее к целевому.
*/

/*
Input: root = [4,2,5,1,3], target = 3.714286
Output: 4
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
 * @param {number} target
 * @return {number}
 */
var closestValue = function(root, target) {
    if (!root) return -1;
    let closest = root.val; // closest - ближайшее

    const dfs = (node) => {
        if (node === null) return;
        if (Math.abs(target - closest) > Math.abs(target - node.val)) {
            closest = node.val;
        }

        if (target < node.val) dfs(node.left);
        else dfs(node.right);
    }

    dfs(root)
    return closest;
}

const example1 = [4,2,5,1,3];
const tree1 = buildTree(example1);
console.log(closestValue(tree1, 3.714286)); // 4
