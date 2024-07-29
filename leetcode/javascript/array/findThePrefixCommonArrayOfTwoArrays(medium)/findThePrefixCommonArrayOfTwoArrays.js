/* https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/

Вам даны две целочисленные перестановки A и B с нулевым индексом длины n.
Общий префиксный массив A и B — это массив C, такой, что C[i] равен количеству чисел, которые присутствуют в индексе i
или перед ним как в A, так и в B.
Верните общий массив префиксов A и B.
Последовательность из n целых чисел называется перестановкой, если она содержит все целые числа от 1 до n ровно один раз.

Input: A = [1,3,2,4], B = [3,1,2,4]
Output: [0,2,3,4]
Объяснение: При i = 0: нет общих чисел, поэтому C[0] = 0.
При i = 1: 1 и 3 являются общими в A и B, поэтому C[1] = 2.
При i = 2: 1, 2 и 3 являются общими в A и B, поэтому C[2] = 3.
При i = 3: 1, 2, 3 и 4 являются общими в A и B, поэтому C[3] = 4.
*/

/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var findThePrefixCommonArray = function(A, B) {
  const res = [];
  const set = new Set();
  let count = 0;

  for (let i = 0; i < A.length; i++) {
    if (set.has(A[i])) count++;
    if (set.has(B[i])) count++;
    if (A[i] === B[i]) count++;
    set.add(A[i]);
    set.add(B[i]);
    res.push(count);
  }

  return res;
};

console.log(findThePrefixCommonArray([1,3,2,4], [3,1,2,4])); // [0,2,3,4]