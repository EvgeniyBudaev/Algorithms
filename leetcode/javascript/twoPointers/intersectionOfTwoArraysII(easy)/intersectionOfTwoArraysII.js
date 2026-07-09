/* 350. Intersection of Two Arrays II
https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

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
var intersect = function (nums1, nums2) {
    // Сортируем оба массива по возрастанию.
    // В JS sort() по умолчанию сортирует как строки, поэтому нужна функция сравнения.
    nums1.sort((a, b) => a - b);
    nums2.sort((a, b) => a - b);

    let left = 0;
    let right = 0;
    const result = [];

    // Проходим по обоим массивам с помощью двух указателей
    while (left < nums1.length && right < nums2.length) {
        if (nums1[left] < nums2[right]) {
            left++;
        } else if (nums1[left] > nums2[right]) {
            right++;
        } else {
            // Если элементы равны, добавляем в результат и двигаем оба указателя
            result.push(nums1[left]);
            left++;
            right++;
        }
    }

    return result;
};