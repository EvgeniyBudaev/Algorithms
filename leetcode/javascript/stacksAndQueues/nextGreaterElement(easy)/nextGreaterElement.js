/*
Следующий больший элемент некоторого элемента x в массиве — это первый больший элемент, находящийся справа от x в том же
массиве.
Вам даны два различных целочисленных массива с нулевым индексом nums1 и nums2, где nums1 — это подмножество nums2.
Для каждого 0 <= i < nums1.length найдите индекс j такой, что nums1[i] == nums2[j] и определите следующий больший
элемент nums2[j] в nums2. Если следующего большего элемента нет, то ответом на этот запрос будет -1.
Верните массив ans длины nums1.length такой, что ans[i] является следующим большим элементом, как описано выше.
*/

/*
Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
 */

/*
Input: nums1 = [2,4], nums2 = [1,2,3,4]
Output: [3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
- 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var nextGreaterElement = function(nums1, nums2) {
    const stack = [];
    const map = {};

    for (let i = nums2.length - 1; i >= 0; --i) {
        while (stack.length && stack[stack.length - 1] <= nums2[i]) {
            stack.pop();
        }
        map[nums2[i]] = stack.length ? stack[stack.length - 1] : -1;
        stack.push(nums2[i]);
    }

    return nums1.map(num => map[num]);
};

console.log(nextGreaterElement([4,1,2], [1,3,4,2])); // [-1,3,-1]