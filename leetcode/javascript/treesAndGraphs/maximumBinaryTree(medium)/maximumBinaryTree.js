/* https://leetcode.com/problems/maximum-binary-tree/description/

Вам дан целочисленный массив чисел без дубликатов. Максимальное двоичное дерево можно построить рекурсивно из чисел,
используя следующий алгоритм:
Создайте корневой узел, значение которого является максимальным значением в числах.
Рекурсивно постройте левое поддерево на префиксе подмассива слева от максимального значения.
Рекурсивно постройте правое поддерево на суффиксе подмассива справа от максимального значения.
Возвращает максимальное двоичное дерево, построенное из чисел.

Input: nums = [3,2,1,6,0,5]
Output: [6,3,5,null,2,0,null,null,1]
Объяснение: Рекурсивные вызовы следующие:
- Самое большое значение в [3,2,1,6,0,5] — 6. Левый префикс — [3,2,1], а правый суффикс — [0,5].
    - Наибольшее значение в [3,2,1] — 3. Левый префикс — [], а правый суффикс — [2,1].
        — Пустой массив, поэтому дочерних элементов нет.
        - Наибольшее значение в [2,1] — 2. Левый префикс — [], а правый суффикс — [1].
            — Пустой массив, поэтому дочерних элементов нет.
            — Только один элемент, поэтому дочерним элементом является узел со значением 1.
    - Максимальное значение в [0,5] — 5. Левый префикс — [0], правый суффикс — [].
        — Только один элемент, поэтому дочерним элементом является узел со значением 0.
        — Пустой массив, поэтому дочерних элементов нет.
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
 * @param {number[]} nums
 * @return {TreeNode}
 */
var constructMaximumBinaryTree = function(nums) {
    let max = Math.max(...nums);
    let root = new TreeNode(max);
    let index = nums.indexOf(max);

    if (nums.slice(0,index).length > 0)
        root.left = constructMaximumBinaryTree(nums.slice(0, index));

    if (nums.slice(index + 1).length > 0)
        root.right = constructMaximumBinaryTree(nums.slice(index + 1));

    return root;
};

console.log(constructMaximumBinaryTree([3,2,1,6,0,5])); // [6,3,5,null,2,0,null,null,1]