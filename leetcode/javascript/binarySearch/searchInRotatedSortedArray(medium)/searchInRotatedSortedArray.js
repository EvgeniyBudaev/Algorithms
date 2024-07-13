/* https://leetcode.com/problems/search-in-rotated-sorted-array/description/

Существует целочисленный массив nums, отсортированный по возрастанию (с разными значениями).

Перед передачей в вашу функцию nums, возможно, поворачивается с неизвестным индексом поворота k (1 <= k < nums.length),
так что результирующий массив имеет вид [nums[k], nums[k+1],... , nums[n-1], nums[0], nums[1], ..., nums[k-1]]
(с индексом 0). Например, [0,1,2,4,5,6,7] можно повернуть с опорным индексом 3 и превратить в [4,5,6,7,0,1,2].

Учитывая массив nums после возможного поворота и целочисленную цель, верните индекс цели, если он находится в nums, или
-1, если он не в nums.

Вы должны написать алгоритм с компиляцией времени выполнения O(log n).

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let low = 0, high = nums.length - 1;

    while (low <= high) {
        let mid = Math.floor((low + high) / 2);
        let guess = nums[mid]; // 7 -> 1 -> 0
        let leftNum = nums[low] , rightNum = nums[high];

        if (guess === target) return mid;

        if (leftNum <= guess) {
            if (leftNum <= target && target < guess) high = mid - 1;
            else low = mid + 1;
        } else {
            if (guess < target && target <= rightNum) low = mid + 1;
            else high = mid - 1;
        }
    }

    return -1;
};

console.log(search([4,5,6,7,0,1,2], 0)); // 4