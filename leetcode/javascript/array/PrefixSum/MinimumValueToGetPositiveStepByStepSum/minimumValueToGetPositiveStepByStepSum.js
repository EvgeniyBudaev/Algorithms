/*
Учитывая массив целых чисел nums, вы начинаете с начального положительного значения startValue.
На каждой итерации вы вычисляете пошаговую сумму startValue плюс элементы в числах (слева направо).
Возвращает минимальное положительное значение startValue, чтобы пошаговая сумма никогда не была меньше 1.
 */

/*
Input: nums = [-3,2,-3,4,2]
Output: 5
 */

/*
Input: nums = [1,2]
Output: 1
Explanation: Minimum start value should be positive.
 */

/*
Input: nums = [1,-2,-3]
Output: 5

 */

/**
 * @param {number[]} nums
 * @return {number}
 */
var minStartValue = function(nums) {
    let startValue = 0; // сумма начального значения
    let minVal = 0;  //минимальное значение

    for (let i = 0; i < nums.length; i++) {
        startValue += nums[i]; // [-3,2,-3,4,2] -> [-3,-1,-4,0,2]
        minVal = Math.min(minVal, startValue);
    }

    return Math.abs(minVal) + 1;
};

console.log(minStartValue([-3,2,-3,4,2])); // 5