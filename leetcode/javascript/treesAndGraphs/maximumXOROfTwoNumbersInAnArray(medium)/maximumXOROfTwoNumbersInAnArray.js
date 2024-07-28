/* https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/description/
solution https://leetcode.com/problems/longest-word-in-dictionary/solutions/5484506/js-rust-trie-approach-w-explanation/

Учитывая целочисленный массив nums, верните максимальный результат операции nums[i] XOR nums[j], где 0 <= i <= j < n.

Input: nums = [3,10,5,25,2,8]
Output: 28
Пояснение: Максимальный результат — 5 XOR 25 = 28.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaximumXOR = function(nums) {
    let L = Math.max(...nums);
    L = L.toString(2).length;
    let maxXor = -Infinity;
    let currXor = 0;
    let prefixes = new Set();

    for(let i = L-1; i >= 0; i--) {
        maxXor <<= 1;
        currXor = maxXor | 1;
        prefixes.clear();

        for(let num of nums) {
            let prefix = num >> i;
            if(prefixes.has(currXor^prefix)) {
                maxXor = currXor;
                break;
            }

            prefixes.add(num >> i);
        }
    }

    return maxXor;
}

console.log(findMaximumXOR([3,10,5,25,2,8])); // 28