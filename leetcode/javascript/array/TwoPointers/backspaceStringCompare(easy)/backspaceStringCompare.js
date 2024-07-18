/* https://leetcode.com/problems/backspace-string-compare/description/

Учитывая две строки s и t, верните true, если они равны, когда обе вводятся в пустые текстовые редакторы. '#' означает
символ возврата.
Обратите внимание, что после возврата пустого текста текст останется пустым.

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
*/

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var backspaceCompare = function(s, t) {
    // Two pointers
    let i = s.length - 1;
    let j = t.length - 1;
    let backspaceCountS = 0;
    let backspaceCountT = 0;

    while (i >= 0 || j >= 0) {
        // Найте позицию следующего допустимого символа в s после учета обратных пробелов.
        while (i >= 0 && (s[i] === '#' || backspaceCountS > 0)) {
            if (s[i] === '#') backspaceCountS++;
            else backspaceCountS--;
            i--;
        }

        // Найте позицию следующего допустимого символа в t после учета обратных пробелов.
        while (j >= 0 && (t[j] === '#' || backspaceCountT > 0)) {
            if (t[j] === '#') backspaceCountT++;
            else backspaceCountT--;
            j--;
        }

        // В обеих строках больше нет допустимых символов, они совпадают.
        if (i < 0 && j < 0) return true;

        // Если допустимые символы в текущих позициях не совпадают, строки не равны.
        if (i < 0 || j < 0 || s[i] !== t[j]) return false;
        i--;
        j--;
    }

    // Если обе строки пусты или содержат только символы возврата, они считаются равными.
    return true;
};

var backspaceCompareWithStack = function(s, t) {
    // Stack
    let build = str => {
        const stack = [];
        for (const c of str) {
            if (c !== '#') {
                stack.push(c);
            } else if (stack.length) {
                stack.pop();
            }
        }
        return stack.join('');
    }
    return build(s) === build(t);
};

console.log(backspaceCompare("ab#c", "ad#c")); // true