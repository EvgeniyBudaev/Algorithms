/* 254. Factor Combinations

Числа можно рассматривать как произведение своих множителей. Например,
8 = 2 х 2 х 2;
  = 2 х 4.
Напишите функцию, которая принимает целое число n и возвращает все возможные комбинации его множителей.

Input: 1
Output: []

Input: 37
Output:[]

Input: 12
Output: [[2, 6], 2, 2, 3], [3, 4]]
*/

/**
 * @param {number} n
 * @return {number[][]}
 */
const getFactors = function (n) {
    const result = [];
    const aux = (remain, start = 2, current = []) => {
        if (remain === 1) {
            if (current.length > 1) result.push([...current]);
            return;
        }
        for (let i = start; i <= remain; i++) {
            if (remain % i === 0) {
                current.push(i);
                aux(remain / i, i, current);
                current.pop();
            }
        }
    }
    aux(n);
    return result;
}

console.log(getFactors(12)); // [[2, 6], 2, 2, 3], [3, 4]]