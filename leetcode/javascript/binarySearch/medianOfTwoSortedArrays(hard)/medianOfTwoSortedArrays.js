/* https://leetcode.com/problems/median-of-two-sorted-arrays/description/
solution https://leetcode.com/problems/median-of-two-sorted-arrays/solutions/4070371/94-96-binary-search-two-pointers/

Учитывая два отсортированных массива nums1 и nums2 размера m и n соответственно, верните медиану двух отсортированных
массивов.
Общая сложность времени выполнения должна составлять O(log (m+n)).

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
*/

// Code Binary Search
/**
 * Time O(log(min(m,n))), Space O(1)
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    if (nums1.length > nums2.length) {
        [nums1, nums2] = [nums2, nums1];
    }

    const m = nums1.length;
    const n = nums2.length;
    let low = 0, high = m; // low: 0, high: 1

    while (low <= high) {
        // partitionX Это середина первого массива. Приблизительно для определения, сколько элементов из первого массива
        // должно быть включено в итоговый набор
        const partitionX = Math.floor((low + high) / 2); // 0 -> 1
        // partitionY позволяет определить, сколько элементов из второго массива должно быть включено в итоговый набор,
        // учитывая уже выбранные элементы из первого массива
        const partitionY = Math.floor((m + n + 1) / 2) - partitionX; // 2 -> 1

        const maxX = (partitionX === 0) ? Number.MIN_SAFE_INTEGER : nums1[partitionX - 1]; // -9007199254740991 -> 2
        const maxY = (partitionY === 0) ? Number.MIN_SAFE_INTEGER : nums2[partitionY - 1]; // 3 -> 1

        const minX = (partitionX === m) ? Number.MAX_SAFE_INTEGER : nums1[partitionX]; // 2 -> 9007199254740991
        const minY = (partitionY === n) ? Number.MAX_SAFE_INTEGER : nums2[partitionY]; // 9007199254740991 -> 3

        if (maxX <= minY && maxY <= minX) {
            if ((m + n) % 2 === 0) {
                return (Math.max(maxX, maxY) + Math.min(minX, minY)) / 2;
            } else {
                return Math.max(maxX, maxY);
            }
        } else if (maxX > minY) {
            high = partitionX - 1;
        } else {
            low = partitionX + 1;
        }
    }

    throw new Error("Input arrays are not sorted.");
};

// Code Two Pointers
/**
 * Time O(m+n), Space O(m+n)
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArraysTwoPointers = function(nums1, nums2) {
    let merged = [];
    let i = 0, j = 0;

    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] < nums2[j]) {
            merged.push(nums1[i++]);
        } else {
            merged.push(nums2[j++]);
        }
    }

    while (i < nums1.length) merged.push(nums1[i++]);
    while (j < nums2.length) merged.push(nums2[j++]);

    let mid = Math.floor(merged.length / 2);
    if (merged.length % 2 === 0) {
        return (merged[mid - 1] + merged[mid]) / 2;
    } else {
        return merged[mid];
    }
};

console.log(findMedianSortedArrays([1,3], [2])); // 2.00000
