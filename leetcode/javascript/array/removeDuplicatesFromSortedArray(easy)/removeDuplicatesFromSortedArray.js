/* https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
Учитывая целочисленный массив чисел, отсортированный в неубывающем порядке, удалите дубликаты на месте так, чтобы каждый
уникальный элемент появлялся только один раз. Относительный порядок элементов должен оставаться неизменным.
Затем верните количество уникальных элементов в числах.
Считайте, что количество уникальных элементов чисел равно k. Чтобы вас приняли, вам нужно сделать следующее:
Измените массив nums так, чтобы первые k элементов nums содержали уникальные элементы в том порядке, в котором они
присутствовали в nums изначально. Остальные элементы nums не важны, как и размер nums.
Вернуть К.

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var removeDuplicates = function(nums) {
    if (nums.length === 0) return [];
    // Инициализируем количество уникальных элементов равным 1
    let k = 1;
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] !== nums[k - 1]) {
            nums[k] = nums[i]; // перезаписываем следующий уникальный элемент
            k++;
        }
    }
    // обрезаем массив до его фактической длины после удаления дубликатов
    nums.length = k;
    // возвращаем массив без дубликатов
    return nums;
};

// var removeDuplicates = function(nums) {
//     let i = nums.length - 1;
//     while (i > 0) {
//         if (nums[i] === nums[i - 1]) {
//             nums.splice(i, 1);
//         }
//         i--;
//     }
//     return nums;
// };

console.log(removeDuplicates([0,0,1,1,1,2,2,3,3,4]));