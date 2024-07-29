/* https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

Путь в бинарном дереве — это последовательность узлов, в которой каждая пара соседних узлов в последовательности имеет
соединяющее их ребро. Узел может появляться в последовательности не более одного раза. Обратите внимание, что путь не
обязательно должен проходить через корень.
Сумма путей пути — это сумма значений узлов в пути.
Учитывая корень двоичного дерева, верните максимальную сумму путей любого непустого пути.

Input: root = [1,2,3]
Output: 6
Пояснение: Оптимальный путь — 2 -> 1 -> 3 с суммой путей 2 + 1 + 3 = 6.
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
var maxPathSum = function(root) {
  let res = {value: -Infinity};

  function solve(node) {
    if(node=== null) return 0;
    let left = solve(node.left);
    let right = solve(node.right);
    let temp = Math.max(Math.max(left, right) + node.val, node.val);
    let ans = Math.max(temp, node.val + left + right);
    res.value = Math.max(res.value, ans);
    return temp;
  }

  solve(root);
  return res.value;
};

const example1 = [1,2,3];
const tree1 = buildTree(example1);
console.log(maxPathSum(tree1)); // 6