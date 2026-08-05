/* 680. Valid Palindrome II
https://leetcode.com/problems/valid-palindrome-ii/description/

Учитывая строку s, верните true, если s может быть палиндромом после удаления из нее не более одного символа.

Input: s = "aba"
Output: true

Input: s = "abca"
Output: true
Объяснение: Вы можете удалить символ 'c'.

Input: s = "abc"
Output: false
*/

/**
 * validPalindrome проверяет, можно ли сделать строку палиндромом, удалив один символ.
 * time: O(n), space: O(1)
 *
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function (s) {
    let left = 0;
    let right = s.length - 1;

    while (left < right) { // Сравниваем символы с обоих концов строки.
        // Если символы не равны, проверяем, можно ли сделать строку палиндромом, удалив один символ.
        if (s[left] !== s[right]) {
            return isPalindrome(s, left + 1, right) || isPalindrome(s, left, right - 1);
        }
        left++;
        right--;
    }

    return true; // Если цикл завершился успешно, строка является палиндромом.
};

// isPalindrome проверяет, является ли подстрока палиндромом.
// time: O(n), space: O(1)
const isPalindrome = function (s, left, right) {
    while (left < right) {
        // Если символы не равны, строка не является палиндромом.
        if (s[left] !== s[right]) {
            return false;
        }
        left++;
        right--;
    }
    return true; // Если цикл завершился успешно, строка является палиндромом.
};