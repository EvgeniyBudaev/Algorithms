/* https://leetcode.com/problems/merge-sorted-array/description/

Вам даны два целочисленных массива nums1 и nums2, отсортированные в порядке неубывания, и два целых числа m и n,
представляющие количество элементов в nums1 и nums2 соответственно.
Объедините nums1 и nums2 в один массив, отсортированный в неубывающем порядке.
Окончательно отсортированный массив не должен возвращаться функцией, а должен храниться внутри массива nums1.
Чтобы учесть это, nums1 имеет длину m + n, где первые m элементов обозначают элементы, которые следует объединить,
а последние n элементов имеют значение 0 и их следует игнорировать. nums2 имеет длину n.

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Пояснение: Мы объединяем массивы [1,2,3] и [2,5,6].
Результатом слияния будет [1,2,2,3,5,6] с подчеркнутыми элементами, взятыми из nums1.

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Пояснение: Мы объединяем массивы [1] и [].
Результатом слияния является [1].

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Пояснение: Мы объединяем массивы [] и [1].
Результатом слияния является [1].
Обратите внимание: поскольку m = 0, в nums1 нет элементов. 0 нужен только для того, чтобы результат слияния мог
поместиться в nums1.
*/

/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    let index1 = m - 1;
    let index2 = n - 1;

    for (let i = m + n - 1; i >= 0; i--) {
        if (index2 < 0) break;
        if (nums1[index1] > nums2[index2]) {
            nums1[i] = nums1[index1];
            index1--;
        } else {
            nums1[i] = nums2[index2];
            index2--;
        }
    }

    console.log(nums1);
};

// var merge = function(nums1, m, nums2, n) {
//     for (let i = m, j = 0; i < m + n; i++, j++) {
//         nums1[i] = nums2[j];
//     }
//     nums1.sort((a, b) => a - b);
//     console.log(nums1);
// };

merge([1,2,3,0,0,0], 3, [2,5,6], 3);
