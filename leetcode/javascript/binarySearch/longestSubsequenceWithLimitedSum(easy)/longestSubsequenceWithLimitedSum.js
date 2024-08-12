/* https://leetcode.com/problems/longest-subsequence-with-limited-sum/description/

Вам дан целочисленный массив nums длины n и запросы целочисленного массива queries длиной m.
Верните answer в виде массива длиной m, где answer[i] — максимальный размер подпоследовательности, которую можно взять
из чисел, чтобы сумма ее элементов была меньше или равна запросу [i].
Подпоследовательность — это массив, который можно получить из другого массива путем удаления некоторых элементов или их
отсутствия без изменения порядка остальных элементов.

Input: nums = [4,5,2,1], queries = [3,10,21]
Output: [2,3,4]
Пояснение: На запросы мы отвечаем следующим образом:
- Сумма подпоследовательности [2,1] меньше или равна 3. Можно доказать, что 2 — максимальный размер такой
подпоследовательности, поэтому ответ[0] = 2.
- Сумма подпоследовательности [4,5,1] меньше или равна 10. Можно доказать, что 3 — максимальный размер такой
подпоследовательности, поэтому ответ[1] = 3.
- Сумма подпоследовательности [4,5,2,1] меньше или равна 21. Можно доказать, что 4 — это максимальный размер такой
подпоследовательности, поэтому ответ[2] = 4. 4.

Input: nums = [2,3,4,5], queries = [1]
Output: [0]
Объяснение: Пустая подпоследовательность — единственная подпоследовательность, сумма которой меньше или равна 1,
поэтому ответ[0] = 0.
 */

/**
 * @param {number[]} nums
 * @param {number[]} queries
 * @return {number[]}
 */
var answerQueries = function(nums, queries) {
    nums.sort((a, b) => a - b); // [1, 2, 4, 5]
    const ans = [];
    for (let query of queries) {
        let count = 0, sum = 0;
        for  (let i = 0; i < nums.length; i++) {
            if (sum + nums[i] <= query) {
                sum += nums[i];
                count++;
            }
        }
        ans.push(count);
    }
    return ans;
};

var answerQueriesV2 = function(nums, queries) {
    nums.sort((a,b) => a - b); // [1, 2, 4, 5]
    const prefixSum = [nums[0]];

    for (let i = 1; i < nums.length; i++) {
        prefixSum.push(nums[i] + prefixSum[i - 1]);
    }

    const answer = [];

    for (let i = 0; i < queries.length; i++) {
        const target = queries[i];
        let [left, right] = [0, prefixSum.length - 1]; // prefixSum: [1, 3, 7, 12]

        while(left <= right){
            let mid = Math.floor((left + right) / 2);
            if (prefixSum[mid] > target) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        answer.push(left);
    }

    return answer;
};

console.log(answerQueries([4,5,2,1], [3,10,21])); // [2,3,4]