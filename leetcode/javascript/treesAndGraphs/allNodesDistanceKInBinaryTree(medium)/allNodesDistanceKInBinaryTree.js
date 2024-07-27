/*
Учитывая корень двоичного дерева, значение целевого узла и целое число k, верните массив значений всех узлов,
находящихся на расстоянии k от целевого узла. Вы можете вернуть ответ в любом порядке.

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
Output: [7,4,1]
Пояснение: Узлы, находящиеся на расстоянии 2 от целевого узла (со значением 5), имеют значения 7, 4 и 1.

Input: root = [1], target = 1, k = 3
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
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function distanceK(root, target, k) {
    if (!root) return [0];

    const g = {};

    const dfs = (node, parent = null) => {
        if (!node) return;

        g[node.val] ??= [];
        if (parent) g[node.val].push(parent.val);
        if (node.left) g[node.val].push(node.left.val);
        if (node.right) g[node.val].push(node.right.val);

        dfs(node.left, node);
        dfs(node.right, node);
    };

    dfs(root);

    const vis = new Set();
    let q = [target.val];

    while (q.length) {
        if (!k--) return q;

        const nextQ = [];

        for (const x of q) {
            if (vis.has(x)) continue;

            vis.add(x);
            nextQ.push(...g[x].filter(x => !vis.has(x)));
        }

        q = nextQ;
    }

    return [];
}