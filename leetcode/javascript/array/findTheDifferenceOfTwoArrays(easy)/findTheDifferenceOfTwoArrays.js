/*

Учитывая два целочисленных массива nums1 и nums2 с индексом 0, верните ответ в виде списка размером 2, где:

ответ[0] — это список всех различных целых чисел в nums1, которых нет в nums2.
ответ[1] — это список всех различных целых чисел в nums2, которых нет в nums1.
Обратите внимание, что целые числа в списках могут возвращаться в любом порядке.

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Объяснение:
Для nums1 nums1[1] = 2 присутствует в индексе 0 в nums2, тогда как nums1[0] = 1 и nums1[2] = 3 отсутствуют в nums2.
Следовательно, ответ[0] = [1,3].
Для nums2 nums2[0] = 2 присутствует в индексе 1 nums1, тогда как nums2[1] = 4 и nums2[2] = 6 отсутствуют в nums2.
Следовательно, ответ[1] = [4,6].
*/

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[][]}
 */
var findDifference = function(nums1, nums2) {
    let s1 = new Set(nums1);
    let s2 = new Set(nums2);
    let ans = [[], []];

    for (let i of s1) {
        if (!s2.has(i)) ans[0].push(i);
    }

    for (let i of s2) {
        if (!s1.has(i)) ans[1].push(i);
    }

    return ans;
};

const nums1 = [1,2,3];
const nums2 = [2,4,6];
console.log(findDifference(nums1, nums2)); // [[1,3],[4,6]]