/* 125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/description/

Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.
Учитывая строку s, верните true, если это палиндром, или false в противном случае.

Input: s = "A man, a plan, a canal: Panama"
Output: true
Объяснение: "amanaplanacanalpanama" палиндром.

Input: s = "race a car"
Output: false
Объяснение: "raceacar" не палиндром.

Input: s = " "
Output: true
Объяснение: s — это пустая строка "" после удаления небуквенно-цифровых символов.
Поскольку пустая строка читается одинаково и вперед, и назад, она является палиндромом.
*/

/**
 * isPalindrome проверяет, является ли строка s палиндромом после нормализации.
 * time: O(n), где n - количество символов в строке, space: O(1)
 *
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
    let left = 0;
    let right = s.length - 1;

    while (left < right) {
        const skipLeft = !isAlphaNumeric(s[left]); // пропускаем символы, которые не являются буквенно-цифровыми
        if (skipLeft) {
            left++;
            continue;
        }

        const skipRight = !isAlphaNumeric(s[right]); // пропускаем символы, которые не являются буквенно-цифровыми
        if (skipRight) {
            right--;
            continue;
        }

        // сравниваем символы, которые являются буквенно-цифровыми
        const endsEqual = s[left].toLowerCase() === s[right].toLowerCase();
        if (!endsEqual) {
            return false;
        }

        left++;
        right--;
    }

    return true; // строка является палиндромом
};

// isAlphaNumeric проверяет, является ли переданный символ буквенно-цифровым
// time: O(1), space: O(1)
const isAlphaNumeric = (c) => {
    return /[a-zA-Z0-9]/.test(c);
};