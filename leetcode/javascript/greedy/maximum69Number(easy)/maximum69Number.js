/*
Вам дано целое положительное число, состоящее только из цифр 6 и 9.
Верните максимальное число, которое вы можете получить, изменив не более одной цифры
(6 становится 9, а 9 становится 6).
*/

/*
Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.
 */

/*
Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.
 */

/*
Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.
 */

/**
 * @param {number} num
 * @return {number}
 */
var maximum69Number  = function(num) {
    let digit = num.toString().split('');

    for (let i = 0; i < digit.length; i++) {
        if (digit[i] === '6') {
            digit[i] = '9';
            break;
        }
    }

    return Number(digit.join(''));
};

console.log(maximum69Number(9669)); // 9969