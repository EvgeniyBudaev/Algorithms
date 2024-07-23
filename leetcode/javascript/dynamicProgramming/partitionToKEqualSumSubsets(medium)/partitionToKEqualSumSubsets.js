/* https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/

Учитывая целочисленный массив nums и целое число k, верните true, если можно разделить этот массив на k непустые
подмножества, все суммы которых равны.

Input: nums = [4,3,2,3,5,2,1], k = 4
Output: true
Explanation: It is possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal

Input: nums = [1,2,3,4], k = 3
Output: false
*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var canPartitionKSubsets = function(nums, k) {
    const total = nums.reduce((sum, num) => sum + num, 0);
    // сразу возвращайте false, если среди k подмножеств нет равной суммы
    if (total % k !== 0) return false;

    const target = total / k;
    const visited = new Array(nums.length).fill(false);

    const canPartition = (start, numberOfSubsets, currentSum) => {
        // базовый вариант
        if (numberOfSubsets === 1) return true;
        // когда подмножество найдено, мы запускаем еще один поиск, чтобы найти оставшиеся подмножества из непосещенных
        // элементов.
        if (currentSum === target) {
            return canPartition(0, numberOfSubsets - 1, 0);
        }
        for (let i = start; i < nums.length; i++) {
            if (!visited[i]) {
                visited[i] = true;
                // запустите поиск, чтобы найти другие элементы, которые будут соответствовать цели текущего элемента.
                if (canPartition(i + 1, numberOfSubsets, currentSum + nums[i])) return true;
                // сброс, чтобы включить возврат
                visited[i] = false;
            }
        }
        return false;
    };

    return canPartition(0, k, 0);
};

const nums = [4,3,2,3,5,2,1];
const k = 4;
console.log(canPartitionKSubsets(nums, k)); // true