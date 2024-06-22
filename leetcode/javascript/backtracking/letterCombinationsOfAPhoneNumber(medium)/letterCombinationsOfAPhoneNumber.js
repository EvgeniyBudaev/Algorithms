// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

/*
Учитывая строку, содержащую цифры от 2 до 9 включительно, верните все возможные комбинации букв, которые может
представлять число. Верните ответ в любом порядке.
Соответствие цифр буквам (как на кнопках телефона) приведено ниже. Обратите внимание, что 1 не соответствует
никаким буквам.
*/

/*
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 */

/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function(digits) {
    if (!digits.length) return [];

    const digitToLetters = {
        '2': 'abc',
        '3': 'def',
        '4': 'ghi',
        '5': 'jkl',
        '6': 'mno',
        '7': 'pqrs',
        '8': 'tuv',
        '9': 'wxyz'
    };

    const res = [];

    function backtrack(idx, comb) {
        if (idx === digits.length) {
            res.push(comb);
            return;
        }

        for (const letter of digitToLetters[digits[idx]]) {
            backtrack(idx + 1, comb + letter);
        }
    }

    backtrack(0, "");

    return res;
};

console.log(letterCombinations("23")); // ["ad","ae","af","bd","be","bf","cd","ce","cf"]