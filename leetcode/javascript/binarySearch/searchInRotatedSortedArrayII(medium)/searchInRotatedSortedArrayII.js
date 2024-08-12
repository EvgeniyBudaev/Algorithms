/* https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

Существует целочисленный массив nums, отсортированный в неубывающем порядке (не обязательно с разными значениями).
Перед передачей в вашу функцию nums вращается с неизвестным индексом поворота k (0 <= k < nums.length), так что
результирующий массив имеет вид [nums[k], nums[k+1],..., nums [n-1], nums[0], nums[1], ..., nums[k-1]] (с индексом 0).
Например, [0,1,2,4,4,4,5,6,6,7] можно повернуть с опорным индексом 5 и превратить в [4,5,6,6,7,0,1,2, 4,4].
Учитывая числа массива после вращения и целочисленную цель, верните true, если цель находится в числах, или ложь, если
она не в числах.
Вы должны максимально сократить общее количество этапов операции.

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {boolean}
 */
var search = function(nums, target) {
    return nums.includes(target);
};

var searchBinarySearch = function(nums, target) {
    let low = 0, high = nums.length - 1;

    while (low <= high) {
        let mid = Math.floor((low + high) / 2);
        if (nums[mid] === target) return true;

        if (nums[low] === nums[mid]) {
            low++;
            continue;
        }

        if (nums[low] <= nums[mid]) {
            if (nums[low] <= target && target <= nums[mid]) high = mid - 1;
            else low = mid + 1;
        } else {
            if (nums[mid] <= target && target <= nums[high]) low = mid + 1;
            else high = mid - 1;
        }
    }
    return false;
};

console.log(search([2,5,6,0,0,1,2], 0)); // true