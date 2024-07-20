/* https://leetcode.com/problems/combinations/description/

Учитывая два целых числа n и k, верните все возможные комбинации k чисел, выбранных из диапазона [1, n].
Вы можете вернуть ответ в любом порядке.

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Пояснение: всего есть 4 комбинации на выбор 2 = 6.
Обратите внимание, что комбинации неупорядочены, т. е. [1,2] и [2,1] считаются одной и той же комбинацией.
*/

/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
var combine = function (n, k) {
    let res = [];
    const backtrack = (start = 1, combination = [] ) => {
        if (combination.length === k) {
            res.push([...combination]);
            return;
        }
        for  ( let i = start; i<=n ; i++ ) {
            combination.push(i);
            backtrack(i+1, combination);
            combination.pop();
        }
    }
    backtrack();
    return res;
};

console.log(combine(4, 2)); // [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]