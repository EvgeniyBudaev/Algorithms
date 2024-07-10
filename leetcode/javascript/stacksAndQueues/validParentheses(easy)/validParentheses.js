/* https://leetcode.com/problems/valid-parentheses/description/

Учитывая строку s, содержащую только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка
допустимой.

Входная строка действительна, если:

Открытые скобки должны закрываться скобками того же типа.
Открытые скобки должны закрываться в правильном порядке.
Каждой закрывающей скобке соответствует открытая скобка того же типа.

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false
*/

/**
 * Time O(N) | Space O(N)
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s, stack = []) {
    const map = {
        '}': '{',
        ']': '[',
        ')': '(',
    };

    for (const char of s) {/* Time O(N) */
        const isBracket = (char in map);
        if (!isBracket) { stack.push(char); continue; }/* Space O(N) */

        const isEqual = (stack[stack.length - 1] === map[char]);
        if (isEqual) { stack.pop(); continue; }

        return false;
    }

    return (stack.length === 0);
};

console.log(isValid("()")); // true