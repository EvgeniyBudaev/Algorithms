/* https://leetcode.com/problems/product-of-array-except-self/description/

Учитывая целочисленный массив nums, верните ответ массива, такой, что answer[i] равен произведению всех элементов nums,
кроме nums[i].
Произведение любого префикса или суффикса чисел гарантированно вписывается в 32-битное целое число.
Вы должны написать алгоритм, который работает за время O(n) и не использует операцию деления.

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/

/**
 * Array
 * Time O(N) | Space O(N)
 * https://leetcode.com/problems/product-of-array-except-self/
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function(nums) {
    const result = [];
    let prefix = 1;
    let postfix = 1;

    // Вычисление префиксных произведений
    // В первом цикле for алгоритм проходит по массиву nums слева направо. Для каждого элемента nums[i] значение prefix
    // умножается на текущий элемент, а затем это произведение записывается в соответствующий элемент массива result.
    // Таким образом, result[i] получает значение произведения всех предыдущих элементов массива nums.
    for (let i = 0; i < nums.length; i++) {
        result[i] = prefix;
        prefix *= nums[i];
    }
    // result = [1, 1, 2, 6]
    // prefix = 24

    // Вычисление постфиксных произведений
    // Во втором цикле for алгоритм начинает работу справа налево. Для каждого элемента nums[i] значение postfix
    // умножается на следующий элемент (nums[i + 1]). Затем текущее значение postfix умножается на result[i], что
    // приводит к тому, что result[i] теперь содержит произведение всех элементов после nums[i], кроме самого nums[i].
    // Это значение записывается обратно в result[i].
    for (let i = nums.length - 2; i >= 0; i--) {
        // nums[i] = 3 -> 2 -> 1
        // nums[i + 1] = 4 -> 3 -> 2
        postfix *= nums[i + 1];
        // postfix = 4 -> 12 -> 24
        result[i] *= postfix;
        // result = [1, 1, 8, 6] -> [1, 12, 8, 6] -> [24, 12, 8, 6]
    }

    return result;
};

console.log(productExceptSelf([1,2,3,4])); // [24,12,8,6]