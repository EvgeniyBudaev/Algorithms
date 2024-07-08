/* https://leetcode.com/problems/group-anagrams/description/

Учитывая массив строк strs, сгруппируйте анаграммы вместе. Вы можете вернуть ответ в любом порядке.
Анаграмма — это слово или фраза, образованная путем перестановки букв другого слова или фразы, обычно с использованием
всех исходных букв ровно один раз.

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Input: strs = [""]
Output: [[""]]

Input: strs = ["a"]
Output: [["a"]]
*/

/**
 * @param {string[]} strs
 * @return {string[][]}
 */
// Time complexity: O(n)
// Space complexity: O(n)
var groupAnagrams = function(strs) {
    if (!strs.length) return [[]];
    const seen = {}, arr = [];

    for (let i = 0; i < strs.length; i++) {
        const index = strs[i].split('').sort().join('');
        if (seen[index]) seen[index].push(strs[i]);
        else seen[index] = [strs[i]];
    }
    // seen: { abt: ['bat'], aet: ['eat', 'tea', 'ate'], ant: ['tan', 'nat' }

    for (let i in seen) {
        arr.push(seen[i]);
    }

    return arr;
};

const strs = ["eat","tea","tan","ate","nat","bat"];
console.log(groupAnagrams(strs)); // [["bat"],["nat","tan"],["ate","eat","tea"]]