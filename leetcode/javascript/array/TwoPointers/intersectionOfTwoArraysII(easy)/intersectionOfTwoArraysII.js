/* https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

Учитывая два целочисленных массива nums1 и nums2, верните массив их пересечения. Каждый элемент результата должен
появляться столько раз, сколько он отображается в обоих массивах, и вы можете возвращать результат в любом порядке.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Пояснение: [9,4] также принимается.
*/

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function(nums1, nums2) {
    nums1.sort((a, b) => a - b);
    nums2.sort((a, b) => a - b);

    let x = 0;
    let y = 0;
    const result = [];

    while (x < nums1.length && y < nums2.length) {
        if (nums1[x] < nums2[y]) {
            x++;
        } else if (nums1[x] > nums2[y]) {
            y++;
        } else {
            result.push(nums1[x]);
            x++;
            y++;
        }
    }

    return result;
};

console.log(intersect([1,2,2,1], [2,2])); // [2,2]