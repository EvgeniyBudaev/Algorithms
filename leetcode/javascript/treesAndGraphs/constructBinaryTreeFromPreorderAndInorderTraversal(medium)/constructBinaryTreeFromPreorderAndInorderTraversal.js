/* https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

Учитывая два целочисленных массива preorder и inorder, где preorder — это предварительный обход двоичного дерева, а
inorder — это неупорядоченный обход того же дерева, постройте и верните двоичное дерево.

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Input: preorder = [-1], inorder = [-1]
Output: [-1]
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
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */
var buildTreeMain = function(preorder, inorder) {
    let p = 0, i = 0;
    const build = function (stop) {
        if (inorder[i] !== stop) {
            const root = new TreeNode(preorder[p++]);
            root.left = build(root.val);
            i++;
            root.right = build(stop);
            return root;
        }
        return null;
    }
    return build();
};

const preorder = [3,9,20,15,7];
const inorder = [9,3,15,20,7];
console.log(buildTreeMain(preorder, inorder)); // [3,9,20,null,null,15,7]