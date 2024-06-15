/*
Учитывая корень двоичного дерева, найдите максимальное значение v, для которого существуют разные узлы a и b,
где v = |a.val - b.val| и a является предком b.
Узел a является предком узла b, если либо: любой дочерний узел a равен b, либо любой дочерний узел a является предком
узла b.
*/

/*
Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
 */

/*
Input: root = [1,null,2,null,0,3]
Output: 3
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
var maxAncestorDiff = function(root) {
    let diffOfNodes = (node, max, min)=> {
        /* Когда листовой узел пытается поместить дочерний узел
         * в рекурсию он не сможет
         * найти этот узел и вернуть 0, оставив
         * это вне сравнения (не затрагивая
         * результат).*/
        if(!node) return 0;

        /* Является ли значение текущего узла
         * самый большой или самый маленький среди всех своих
         * непосредственных предков необходимо сравнивать с
         *узнаем, и после сравнения результат
         * обновлено.*/
        max = Math.max(node.val, max);
        min = Math.min(node.val, min);

        let mostDiff = Math.abs(max - min);

        // Запускаем рекурсию
        let leftDiff = diffOfNodes(node.left, max, min);
        let rightDiff = diffOfNodes(node.right, max, min);

        /* Разница текущего узла над
         * предок потенциально не такой большой, как
         * дочерние узлы, поэтому сравните их все три: текущий,
         * Лево и право.*/
        return Math.max(mostDiff, leftDiff, rightDiff);
    };

    return diffOfNodes(root, root.val, root.val)
};

const example1 = [8,3,10,1,6,null,14,null,null,4,7,13];
const tree1 = buildTree(example1);
const example2 = [1,null,2,null,0,3];
const tree2 = buildTree(example2);
console.log(maxAncestorDiff(tree1)); // 7
console.log(maxAncestorDiff(tree2)); // 3
