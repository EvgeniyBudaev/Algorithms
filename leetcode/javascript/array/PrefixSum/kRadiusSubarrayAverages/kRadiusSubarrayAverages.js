/*
Вам дан массив nums с нулевым индексом из n целых чисел и целое число k.

Среднее значение k-радиуса для подмассива чисел с центром в некотором индексе i с радиусом k представляет собой среднее
значение всех элементов в числах между индексами i - k и i + k (включительно). Если до или после индекса i меньше k
элементов, то среднее значение k-радиуса равно -1.

Создайте и верните массив avgs длины n, где avgs[i] — среднее значение k-радиуса для подмассива с центром в индексе i.

Среднее значение элементов x представляет собой сумму элементов x, разделенную на x с использованием целочисленного
деления. Целочисленное деление усекается до нуля, что означает потерю дробной части.

Например, среднее значение четырех элементов 2, 3, 1 и 5 равно (2 + 3 + 1 + 5)/4 = 11/4 = 2,75, что сокращается до 2.
 */

/*
Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
Output: [-1,-1,-1,5,4,4,-1,-1,-1]
 */

/*
Input: nums = [100000], k = 0
Output: [100000]
 */

/*
Input: nums = [8], k = 100000
Output: [-1]
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var getAverages = function(nums, k) {
    const n = nums.length;
    let ans = new Array(n).fill(0); // Initialize ans array with zeros

    let pref = new Array(n).fill(0); // Initialize prefix sum array
    pref[0] = nums[0];
    for (let i = 1; i < n; i++) {
        pref[i] = pref[i - 1] + nums[i];
    }

    for (let i = 0; i < n; i++) {
        let lcnt = i;
        let rcnt = n - i - 1;
        if (lcnt >= k && rcnt >= k) {
            let lsum = 0;
            let rsum = pref[i + k];

            if (i - k > 0) {
                lsum = pref[i - k - 1];
            }

            ans[i] = Math.floor((rsum - lsum) / (2 * k + 1));
        } else {
            ans[i] = -1;
        }
    }

    return ans;
};

console.log(getAverages([7,4,3,9,1,8,5,2,6], 3)); // [-1,-1,-1,5,4,4,-1,-1,-1]