/*
Вам дан корневой узел двоичного дерева поиска (BST) и значение для вставки в дерево.
Верните корневой узел BST после вставки. Гарантируется, что новое значение не существует в исходном BST.
Обратите внимание, что может существовать несколько допустимых способов вставки, если после вставки дерево остается BST.
Вы можете вернуть любой из них.
*/

/*
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
 */

/*
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]
 */

/*
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
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
 * @param {number} val
 * @return {TreeNode}
 */
var insertIntoBST = function(root, val) {
    if (!root) return new TreeNode(val);

    if (root.val < val) {
        // если значение больше корневого значения, выполняем поиск
        // на правой стороне дерева и независимо от поиска
        // возвращает присвоение этого значения справа от корня.
        root.right = insertIntoBST(root.right, val);
    } else if (root.val > val) {
        // если значение больше корневого значения, выполняем поиск
        // в левой части дерева и независимо от поиска
        // возвращает присвоение этого значения слева от корня.
        root.left = insertIntoBST(root.left, val);
    }

    return root;
};

const example1 = [4,2,7,1,3];
const tree1 = buildTree(example1);
console.log(insertIntoBST(tree1, 5)); // [4,2,7,1,3,5]