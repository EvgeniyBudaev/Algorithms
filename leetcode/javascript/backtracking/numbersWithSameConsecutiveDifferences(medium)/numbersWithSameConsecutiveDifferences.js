// https://leetcode.com/problems/numbers-with-same-consecutive-differences/description/

/*
Учитывая два целых числа n и k, верните массив всех целых чисел длины n, где разница между всеми двумя последовательными
цифрами равна k. Вы можете вернуть ответ в любом порядке.
Обратите внимание, что целые числа не должны иметь ведущих нулей. Целые числа 02 и 043 не допускаются.
*/

/*
Input: n = 3, k = 7
Output: [181,292,707,818,929]
Объяснение: Обратите внимание: 070 не является допустимым числом, поскольку в нем есть ведущие нули.
 */

/*
Input: n = 2, k = 1
Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
 */

/**
 * @param {number} n
 * @param {number} k
 * @return {number[]}
 */
var numsSameConsecDiff = function(n, k) {
    let list = [1, 2, 3, 4, 5, 6, 7, 8, 9];

    while (--n > 0) {
        let tmp = [];

        for (let val of list) {
            let rem = val % 10;
            if (rem + k < 10) tmp.push(val * 10 + rem + k);
            if (k !== 0 && rem - k >= 0) tmp.push(val * 10 + rem - k);
        }

        list = tmp;
    }

    return list
};

console.log(numsSameConsecDiff(3, 7)); // [181, 292, 707, 818, 929]