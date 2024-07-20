/* https://leetcode.com/problems/permutations/description/

Учитывая массив различных целых чисел, верните все возможные перестановки. Вы можете вернуть ответ в любом порядке.

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Input: nums = [1]
Output: [[1]]
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function(nums) {
    const ans = [];
    const backTrack = (curr) => {
        if (curr.length === nums.length) {
            ans.push([...curr]);
            return;
        }
        for (let num of nums) {
            if (!curr.includes(num)) {
                curr.push(num);
                backTrack(curr);
                curr.pop();
            }
        }
    }
    backTrack([]);
    return ans;
};

console.log(permute([1,2,3])); // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]