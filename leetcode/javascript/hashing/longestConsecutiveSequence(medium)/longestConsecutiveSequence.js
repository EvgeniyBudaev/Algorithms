/* https://leetcode.com/problems/longest-consecutive-sequence/description/

Учитывая несортированный массив целых чисел nums, верните длину самой длинной последовательной последовательности
элементов.
Вы должны написать алгоритм, который работает за время O(n).

Input: nums = [100,4,200,1,3,2]
Output: 4
Объяснение: Самая длинная последовательная последовательность элементов — [1,2,3,4]. Следовательно, его длина равна 4

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

/**
 * Sort - HeapSort Space O(1) | QuickSort Space O(log(K))
 * Greedy - Max Score
 * Time O (N * log(N)) | Space O(1)
 * @param {number[]} nums
 * @return {number}
 */
var longestConsecutive = function(nums) {
    if (!nums.length) return 0;
    nums.sort((a, b) => a - b); /* Time O(N * log(N)) | Space O(1 || log(N)) */
    return search(nums); /* Time O(N) */
};

const search = (nums) => {
    let maxScore = 1, score = 1;

    for (let i = 1; i < nums.length; i++) { /* Time O(N) */
        const isPrevDuplicate = nums[i - 1] === nums[i];
        if (isPrevDuplicate) continue;

        const isStreak = nums[i] === ((nums[i - 1]) + 1);
        if (isStreak) { score++; continue; }

        maxScore = Math.max(maxScore, score);
        score = 1;
    }

    return Math.max(maxScore, score);
}

console.log(longestConsecutive([100,4,200,1,3,2])); // 4