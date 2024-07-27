/* https://leetcode.com/problems/path-sum-ii/description/

Учитывая корень двоичного дерева и целое число targetSum, верните все пути от корня к листу, где сумма значений узлов в
пути равна targetSum. Каждый путь должен быть возвращен в виде списка значений узлов, а не ссылок на узлы.

Путь от корня к листу — это путь, начинающийся от корня и заканчивающийся любым листовым узлом. Лист — это узел, не
имеющий дочерних элементов.

Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Объяснение: Есть два пути, сумма которых равна targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22
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
 * @return {number[][]}
 */
var pathSum = function(root, targetSum) {
    let r = [];

    function traverse(node, path, sum){
        if(!node) return;
        sum += node.val;
        path.push(node.val);

        if (!node.left && !node.right) {
            if(sum === targetSum) r.push(path.slice());
        }

        traverse(node.left, path, sum);
        traverse(node.right, path, sum);
        path.pop();
    }

    traverse(root, [], 0);
    return r;
};

const example1 = [5,4,8,11,null,13,4,7,2,null,null,5,1];
const tree1 = buildTree(example1);
console.log(pathSum(tree1, 22)); // [[5,4,11,2],[5,8,4,5]]