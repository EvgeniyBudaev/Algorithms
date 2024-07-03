/* https://leetcode.com/problems/number-of-ways-to-split-array/description/
Дан массив чисел nums, найти количество способов разбиения массива на 2 части, таким образом, чтобы первая часть имела
сумму элементов больше или равна сумме элементов второй части массива.
Вторая часть массива должна иметь как минимум 1 элемент.

Input: nums = [4, 2, 1, 1]
Output: 3

Input: nums = [10,4,-8,7]
Output: 2
 */

/**
 * @param {number[]} nums
 * @return {number}
 */
// Подсчет количества способов разбиения массива на 2 части. Сложность алгоритма O(n), по памяти O(1).
var waysToSplitArray = function(nums) {
    let answer =  0; // answer - количество способов разбиения массива на  2 части.
    let leftSection = 0;  // leftSection - сумма элементов первой части массива.
    let total  =  0; // total - сумма элементов.

    for (let i = 0; i < nums.length; i++) {
        total += nums[i]; // 8
    }

    // [4, 2, 1, 1]
    for (let i = 0; i < nums.length - 1; i++) {
        leftSection += nums[i];
        const rightSection = total - leftSection; // rightSection  - сумма элементов второй части массива.

        if (leftSection >= rightSection) {
            answer++;
        }
    }

    return answer;
};

// Подсчет количества способов разбиения массива на 2 части. Сложность алгоритма O(n), по памяти O(n).
// var waysToSplitArray = function(nums) {
//     const prefix = [nums[0]];
//
//     // Формируем массив prefix sum [4, 6, 7, 8]
//     for (let i = 1; i < nums.length; i++) {
//         prefix.push(nums[i] + prefix[i - 1]);
//     }
//
//     let answer =  0; // answer - количество способов разбиения массива на  2 части.
//
//     for (let i = 0; i < nums.length - 1; i++) {
//         let leftSection = prefix[i];  // leftSection  - сумма элементов первой части массива.
//         let rightSection = prefix[nums.length - 1] - prefix[i]; // rightSection  - сумма элементов второй части массива.
//
//         if (leftSection >= rightSection) {
//             answer++;
//         }
//     }
//
//     return answer;
// };

console.log(waysToSplitArray([4, 2, 1, 1])); // 3