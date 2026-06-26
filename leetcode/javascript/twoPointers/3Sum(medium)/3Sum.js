/* 15. 3Sum
https://leetcode.com/problems/3sum/description/

Учитывая целочисленный массив nums, верните все тройки [nums[i], nums[j], nums[k]] такие,
что i != j, i != k и j != k и nums[i] + nums[j] + nums[к] == 0.
Обратите внимание, что набор решений не должен содержать повторяющихся троек.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
Отдельные тройки — это [-1,0,1] и [-1,-1,2].
Обратите внимание, что порядок вывода и порядок троек не имеют значения.

Input: nums = [0,1,1]
Output: []
Пояснение: Единственная возможная тройка не дает в сумме 0.

Input: nums = [0,0,0]
Output: [[0,0,0]]
Пояснение: Сумма единственно возможной тройки равна 0.
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
    const res = []; // Массив для хранения результатов
    // Сортируем массив по возрастанию. 
    // В JS sort() по умолчанию сортирует как строки, поэтому нужна функция сравнения (a, b) => a - b
    nums.sort((a, b) => a - b);

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];

        // Если текущее число больше 0, дальнейшие числа тоже будут больше 0, и сумма не может быть 0
        if (num > 0) {
            break;
        }

        // Пропускаем дубликаты для первого элемента тройки
        if (i > 0 && num === nums[i - 1]) {
            continue;
        }

        let left = i + 1;            // Указатель на следующий элемент после текущего числа
        let right = nums.length - 1; // Указатель на последний элемент

        while (left < right) {
            const sum = num + nums[left] + nums[right];

            if (sum > 0) {
                right--; // Сумма слишком большая, уменьшаем правый указатель
            } else if (sum < 0) {
                left++; // Сумма слишком маленькая, увеличиваем левый указатель
            } else {
                // Нашли тройку
                res.push([num, nums[left], nums[right]]);
                left++;
                right--;

                // Пропускаем дубликаты для левого и правого указателей
                while (left < right && nums[left] === nums[left - 1]) {
                    left++;
                }
                while (left < right && nums[right] === nums[right + 1]) {
                    right--;
                }
            }
        }
    }

    return res;
};