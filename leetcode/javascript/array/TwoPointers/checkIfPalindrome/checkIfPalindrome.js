/*
Учитывая строку s, верните true, если это палиндром, и false в противном случае.
Строка является палиндромом, если она читается как вперед, так и назад одинаково.
 */

/**
 * @param {string} s
 * @return {boolean}
 */
var checkIfPalindrome = function(s) {
    let left = 0;
    let right = s.length - 1;

    while (left < right) {
        if (s[left] !== s[right]) {
            return false;
        }

        left++;
        right--;
    }

    return true;
}

console.log(checkIfPalindrome("racecar"));
console.log(checkIfPalindrome("aleba"));