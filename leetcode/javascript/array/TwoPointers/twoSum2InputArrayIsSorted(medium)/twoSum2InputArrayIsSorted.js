/* https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

Учитывая массив целых чисел с индексом 1, который уже отсортирован в неубывающем порядке, найдите два числа, сумма
которых дает определенное целевое число. Пусть эти два числа будут числами[индекс1] и числами[индекс2],
где 1 <= index1 < index2 <= numbers.length.

Верните индексы двух чисел, индекс1 и индекс2, добавленные на единицу, в виде целочисленного массива
[индекс1, индекс2] длины 2.

Тесты генерируются так, что существует ровно одно решение. Вы не можете использовать один и тот же элемент дважды.
Ваше решение должно использовать только постоянное дополнительное пространство.

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Пояснение: Сумма 2 и 7 равна 9. Следовательно, index1 = 1, index2 = 2. Мы возвращаем [1, 2].

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
 */

/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    let left = 0, right = numbers.length - 1;

    while (left < right) {
        const sum = numbers[left] + numbers[right];
        if (sum === target) return [left + 1, right + 1];
        else if (sum < target) left++;
        else right--;
    }

    return [-1, -1];
};

console.log(twoSum([2,7,11,15], 9));