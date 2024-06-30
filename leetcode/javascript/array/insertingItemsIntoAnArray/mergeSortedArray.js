/*
Вам даны два целочисленных массива nums1 и nums2, отсортированные в порядке неубывания, и два целых числа m и n,
представляющее количество элементов в nums1 и nums2 соответственно.

Объедините nums1 и nums2 в один массив, отсортированный в неубывающем порядке.

Окончательно отсортированный массив не должен возвращаться функцией, а должен храниться внутри массива nums1.
Для этого nums1 имеет длину m + n, где первые m элементов обозначают элементы, которые должны
быть объединены, а последние n элементов установлены в 0 и их следует игнорировать. nums2 имеет длину n.
*/

/*
Example 1:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

Example 2:
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]

Example 3:
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]

Constraints:
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109
 */

/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
// var merge = function(nums1, m, nums2, n) {
//     for (let i = m, j = 0; i < m + n; i++, j++) {
//         nums1[i] = nums2[j];
//     }
//     nums1.sort((a, b) => a - b);
//     console.log(nums1);
// };

 const merge = function(nums1, m, nums2, n) {
    let p1 = m - 1;
    let p2 = n - 1;

    for (let i = m + n - 1; i >= 0; i--) {
        if (p2 < 0) break;

        if (nums1[p1] > nums2[p2]) {
            nums1[i] = nums1[p1];
            p1--;
        } else {
            nums1[i] = nums2[p2];
            p2--;
        }
    }

    console.log(nums1);
};

merge([1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3);
