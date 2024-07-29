/* https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/

Учитывая строку s, верните максимальное количество уникальных подстрок, на которые можно разбить данную строку.
Вы можете разделить строку на любой список непустых подстрок, где объединение подстрок образует исходную строку.
Однако вам необходимо разделить подстроки так, чтобы все они были уникальными.
Подстрока — это непрерывная последовательность символов внутри строки.

Input: s = "ababccc"
Output: 5
Пояснение: Один из способов максимального разделения — это ['a', 'b', 'ab', 'c', 'cc'].
Разделение типа ['a', 'b', 'a', 'b', 'c', 'cc'] недопустимо, поскольку у вас есть 'a' и 'b' несколько раз.

Input: s = "aba"
Output: 2
Пояснение: Один из способов максимального разделения — ['a', 'ba'].

Input: s = "aa"
Output: 1
Объяснение: Дальнейшее разделение строки невозможно.
*/

/**
 * @param {string} s
 * @return {number}
 */
var maxUniqueSplit = function(s) {
  return getMax(0, s, new Set());
};

function getMax(start, s, set) {
  if (start === s.length) return set.size;
  let count = 0;
  for (let i = start + 1; i <= s.length; i++) {
    let sub = s.substring(start, i);
    if (!set.has(sub)) {
      set.add(sub);
      count = Math.max(count, getMax(i, s, set));
      set.delete(sub);
    }
  }
  return count;
}

console.log(maxUniqueSplit("ababccc"));