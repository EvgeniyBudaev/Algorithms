/* https://leetcode.com/problems/generate-parentheses/description/

Учитывая n пар круглых скобок, напишите функцию, которая генерирует все комбинации правильных круглых скобок.

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]
*/

/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n) {
    const result = [];

    function backtrack(s = '', left = 0, right = 0) {
        if (s.length === 2 * n) {
            result.push(s);
            return;
        }
        if (left < n) backtrack(s + '(', left + 1, right);
        if (right < left) backtrack(s + ')', left, right + 1);
    }

    backtrack();
    return result;
};

console.log(generateParenthesis(3)); // ["((()))","(()())","(())()","()(())","()()()"]