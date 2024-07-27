/* https://leetcode.com/problems/path-sum-iii/description/

Учитывая корень двоичного дерева и целое число targetSum, верните количество путей, в которых сумма значений вдоль пути
равна targetSum.
Путь не обязательно должен начинаться или заканчиваться в корне или листе, но он должен идти вниз (т. е. идти только от
родительских узлов к дочерним узлам).

Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Пояснение: Показаны пути, сумма которых равна 8.
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
 * @param {number} targetSum
 * @return {number}
 */
var pathSum = function(root, targetSum) {
    let totalPaths = 0;
    let hashMap = new Map();

    const dfs = (node, currSum) => {
        if (node == null) return;
        currSum += node.val;
        if (currSum === targetSum) totalPaths++;
        totalPaths += (hashMap.get(currSum - targetSum) || 0);
        hashMap.set(currSum, (hashMap.get(currSum) || 0) + 1);
        dfs(node.left, currSum);
        dfs(node.right, currSum);
        hashMap.set(currSum, hashMap.get(currSum) - 1);
    }

    dfs(root, 0);
    return totalPaths;
};

const example1 = [10,5,-3,3,2,null,11,3,-2,null,1];
const tree1 = buildTree(example1);
console.log(pathSum(tree1, 8)); // 3