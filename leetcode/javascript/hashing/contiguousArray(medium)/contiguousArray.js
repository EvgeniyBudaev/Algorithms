/*
Учитывая числа двоичного массива, верните максимальную длину непрерывного подмассива с равным количеством 0 и 1.
 */

/*
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
 */

/*
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
 */

/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaxLength = function(nums) {
    let maxLen = 0; // Текущая максимальная длина подмассива
    // count - Отслеживает разницу между количеством единиц и нулей в рассматриваемой части массива.
    // Если count положительно, это означает, что больше единиц, чем нулей; если отрицательно — наоборот.
    let count = 0;
    let map = {[0]: -1};

    for (let i = 0; i < nums.length; i++) {
        count += nums[i] === 1 ? 1 : -1;

        // Поиск подмассивов с равным количеством нулей и единиц
        // Если да, то вычисляем длину текущего подмассива как разницу между текущим индексом i и индексом последнего
        // вхождения такого же count (map[count]). Обновляем maxLen, если текущая длина больше предыдущего максимума.
        if (map[count] !== undefined) {
            maxLen = Math.max(maxLen, i - map[count]);
        } else {
            // Если нет, то сохраняем текущий индекс i в map под ключом count.
            map[count] = i;
        }
    }

    // Возвращаем максимальную длину подмассива с равным количеством нулей и единиц
    return maxLen;
};

console.log(findMaxLength([0,1])); // 2