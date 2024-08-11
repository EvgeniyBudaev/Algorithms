/*
Учитывая две строки s и t, верните true, если s является подпоследовательностью t, или false в противном случае.

Подпоследовательность строки — это последовательность символов, которую можно получить путем удаления некоторых
(или ни одного) символов из исходной строки, сохраняя при этом относительный порядок оставшихся символов.
Например, «ace» является подпоследовательностью «abcde», а «aec» — нет.
 */

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
    let left = 0, right = 0;
    while (left < s.length && right < t.length) {
        if (s[left] === t[right]) left++;
        right++;
    }
    return left === s.length;
};

console.log(isSubsequence("ace", "abcde")); // true
console.log(isSubsequence("aec", "abcde")); // false