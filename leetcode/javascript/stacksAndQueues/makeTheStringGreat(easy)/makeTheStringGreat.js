/* https://leetcode.com/problems/make-the-string-great/description/
Дана строка s, состоящая из строчных и прописных английских букв.
Хорошая строка — это строка, в которой нет двух соседних символов s[i] и s[i + 1], где:
0 <= i <= s.length - 2
s[i] — строчная буква, а s[i + 1] — та же буква, но в верхнем регистре или наоборот.
Чтобы сделать строку хорошей, вы можете выбрать два соседних символа, которые делают строку плохой, и удалить их.
Вы можете продолжать делать это, пока строка не станет хорошей.
Верните строку после того, как исправили ее. Ответ гарантированно будет уникальным при заданных ограничениях.
Обратите внимание, что пустая строка тоже подойдет.

Input: s = "leEeetcode"
Output: "leetcode"
Explanation: На первом этапе, либо вы выберете i = 1, либо i = 2, в результате оба варианта «leEeetcode» будут сокращены
до «leetcode».

Input: s = "abBAcC"
Output: ""
Explanation: У нас есть много возможных сценариев, и все они ведут к одному и тому же ответу. Например:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""

Input: s = "s"
Output: "s"
*/

/**
 * @param {string} s
 * @return {string}
 */
var makeGood = function(s) {
    if (!s) return '';
    let stack = [];

    for (let char of s) {
        let top = stack[stack.length - 1];
        if(char !== top && char.toLowerCase() === (top ?? '').toLowerCase()) stack.pop();
        else stack.push(char);
    }

    return stack.join('');
};

console.log(makeGood("leEeetcode")); // "leetcode"