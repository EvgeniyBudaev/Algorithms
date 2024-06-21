/*
Учитывая целочисленный массив nums и целое число k, разбейте числа на k непустые подмассивы так, чтобы наибольшая сумма
любого подмассива была минимизирована.
Верните минимизированную наибольшую сумму разделения.
Подмассив — это непрерывная часть массива.
*/

/*
Input: nums = [7,2,5,10,8], k = 2
Output: 18
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разделить его на [7,2,5] и [10,8], где наибольшая сумма среди двух подмассивов равна всего 18.
 */

/*
Input: nums = [1,2,3,4,5], k = 2
Output: 9
Объяснение: Существует четыре способа разбить числа на два подмассива.
Лучший способ — разбить его на [1,2,3] и [4,5], где наибольшая сумма среди двух подмассивов равна всего 9.
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var splitArray = function(nums, k) {
    let left = Math.max(...nums);
    let right = nums.reduce((a, b) => a + b);
    let ans = -1;

    while (left <= right) {
        let largestSum = Math.floor((left + right) / 2);

        if (countSplitSubarray(largestSum, nums) <= k) {
            ans = largestSum;
            right = largestSum - 1;
        } else {
            left = largestSum + 1;
        }
    }

    return ans;
};

const countSplitSubarray = (largestSum, nums) => {
    let countSubarray = 1;
    let currentSum = 0;

    for (let i = 0; i < nums.length; i++) {
        currentSum += nums[i];

        if (currentSum > largestSum) {
            countSubarray += 1;
            currentSum = nums[i];
        }
    }

    return countSubarray;
};

const nums = [7,2,5,10,8];
console.log(splitArray(nums, 2)); // 18