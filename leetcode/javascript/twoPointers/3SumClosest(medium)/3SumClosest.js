/* 16. 3Sum Closest
https://leetcode.com/problems/3sum-closest/description/

Учитывая целочисленный массив nums длины n и целочисленную цель, найдите три целых числа в nums, сумма которых наиболее
близка к цели. Верните сумму трех целых чисел. Вы можете предположить, что каждый вход будет иметь ровно одно решение.

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Пояснение: Ближайшая к цели сумма равна 2. (-1 + 2 + 1 = 2).

Input: nums = [0,0,0], target = 1
Output: 0
Пояснение: Сумма, ближайшая к целевой, равна 0. (0 + 0 + 0 = 0).
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function (nums, target) {
    // Сортируем массив по возрастанию
    nums.sort((a, b) => a - b);

    const n = nums.length;
    // Инициализируем ближайшую сумму тремя первыми элементами массива
    let closestSum = nums[0] + nums[1] + nums[2];

    // n-2, чтобы гарантировать, что после текущего элемента nums[i] останется 
    // как минимум два элемента для формирования тройки.
    for (let i = 0; i < n - 2; i++) {
        let left = i + 1;
        let right = n - 1;

        while (left < right) {
            const sum = nums[i] + nums[left] + nums[right];

            // Если сумма равна цели, возвращаем её сразу
            if (sum === target) {
                return sum;
            }

            // Обновляем closestSum, если текущая сумма ближе к цели
            if (Math.abs(sum - target) < Math.abs(closestSum - target)) {
                closestSum = sum;
            }

            if (sum < target) {
                // Если сумма меньше цели, двигаем левый указатель вправо
                left++;
            } else {
                // Если сумма больше цели, двигаем правый указатель влево
                right--;
            }
        }
    }

    return closestSum;
};