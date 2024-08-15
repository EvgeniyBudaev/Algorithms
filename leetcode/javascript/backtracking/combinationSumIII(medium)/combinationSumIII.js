/* https://leetcode.com/problems/combination-sum-iii/description/

Найдите все допустимые комбинации k чисел, сумма которых равна n, такие, что выполняются следующие условия:
Используются только цифры от 1 до 9.
Каждое число используется не более одного раза.
Возвращает список всех возможных допустимых комбинаций. Список не должен содержать одну и ту же комбинацию дважды,
комбинации могут возвращаться в любом порядке.

Input: k = 3, n = 7
Output: [[1,2,4]]
Объяснение:
1 + 2 + 4 = 7
Других допустимых комбинаций нет

Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Объяснение:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
Других допустимых комбинаций нет.
*/

/**
 * @param {number} k
 * @param {number} n
 * @return {number[][]}
 */
var combinationSum3 = function(k, n) {
    const res = [];

    const backtracking = (currentDigit, sum, elements) => {
        if (currentDigit > 9) return;
        if (sum > n) return;
        if (elements.length > k) return;

        if (sum === n && elements.length === k) {
            res.push(elements);
            return;
        }

        for (let i = currentDigit + 1; i <= 9; i++) {
            backtracking(i, sum + i, [...elements, i]);
        }
    }

    backtracking(0, 0, []);
    return res;
};

console.log(combinationSum3(3, 7));