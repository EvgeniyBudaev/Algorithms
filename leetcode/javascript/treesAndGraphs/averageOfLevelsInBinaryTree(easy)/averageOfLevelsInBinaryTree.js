/* https://leetcode.com/problems/average-of-levels-in-binary-tree/description/

Учитывая корень двоичного дерева, верните среднее значение узлов на каждом уровне в виде массива. Принимаются ответы в
пределах 10-5 от фактического ответа.

Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Пояснение: Среднее значение узлов на уровне 0 — 3, на уровне 1 — 14,5 и на уровне 2 — 11.
Следовательно, верните [3, 14.5, 11].
*/

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
var averageOfLevels = function(root) {
    const average = [], nodes = [root];

    while (nodes.length) {
        const length = nodes.length;
        let sum = 0;

        for (let idx = 0; idx < length; idx++) {
            const node = nodes.shift();
            sum += node.val;
            if (node.left) nodes.push(node.left);
            if (node.right) nodes.push(node.right);
        }

        average.push(sum / length);
    }

    return average;
};
