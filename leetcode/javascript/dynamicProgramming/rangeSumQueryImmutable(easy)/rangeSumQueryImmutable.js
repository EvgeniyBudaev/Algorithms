/* https://leetcode.com/problems/range-sum-query-immutable/description/

Учитывая целочисленный массив nums, обработайте несколько запросов следующего типа:

Вычислить сумму элементов чисел между индексами слева и справа включительно, где слева <= справа.
Реализуйте класс NumArray:

NumArray(int[] nums) Инициализирует объект целочисленным массивом nums.
int sumRange(int left, int right) Возвращает сумму элементов nums между индексами слева и справа включительно
(т. е. nums[left] + nums[left + 1] + ... + nums[right]).

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Объяснение
NumArray numArray = новый NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // возвращаем (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // возвращаем 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // возвращаем (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
*/

/**
 * @param {number[]} nums
 */
var NumArray = function(nums) {
    this.nums = nums;
};

/**
 * @param {number} left
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function(left, right) {
    let sum = 0;
    for(let i = left; i <= right; i++) sum += this.nums[i];
    return sum;
};

/**
 * Your NumArray object will be instantiated and called as such:
 * var obj = new NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */