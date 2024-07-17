/* https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

Вам дан массив символов-букв, отсортированный в неубывающем порядке, и целевой символ. В буквах есть как минимум два
разных символа.
Верните наименьший буквенный символ, который лексикографически больше целевого. Если такого символа не существует,
вернуть первый символ буквами.

Input: letters = ["c","f","j"], target = "a"
Output: "c"
Пояснение: Самый маленький символ, который лексикографически больше буквы «a», — это «c».
*/

/**
 * @param {string[]} letters
 * @param {string} target
 * @return {string}
 */
var nextGreatestLetter = function(letters, target) {
    const length = letters.length;

    // обрабатывать случай вне массива
    if (target < letters[0]) return letters[0];
    if (target >= letters[length - 1]) return letters[0];

    // простой двоичный поиск для поиска символа
    let low = 0, high = length - 1;
    while (low <= high) {
        const mid = Math.floor((low + high) / 2);
        const guess = letters[mid];
        if (guess > target && target >= letters[mid - 1]) return guess;
        if (guess > target) high = mid - 1;
        else low = mid + 1;
    }
};

console.log(nextGreatestLetter(["c","f","j"], "a")); // "c"